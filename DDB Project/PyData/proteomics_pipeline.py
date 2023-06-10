from Bio import ExPASy
from Bio import SeqIO
import concurrent.futures
import pandas
import argparse
import os
import pyopenms
import json
from PeptideIdentification import PeptideIdentification

def Get_Protein_Fasta_By_Id(id, output_dir=None):
    with ExPASy.get_sprot_raw(id) as handle:
        seq_record = SeqIO.read(handle, "swiss")
    file_name = seq_record.name
    file_name += ".fasta"
    if output_dir != None:
        file_path = os.path.join(output_dir, file_name)
    else:
        file_path = file_name    
    return SeqIO.write(seq_record, file_path, "fasta")


def Convert_Spectrum_To_Dict(msexperiment_obj, lower_index):
    print("Start Converting from MSExperiment object to python dictionary...")

    list_of_spectra = []
    index = lower_index
    for spectrum in msexperiment_obj.getSpectra():
        mzs = [float(mz) for mz in spectrum.get_peaks()[0]]
        Is  = [float(i) for i in spectrum.get_peaks()[1]]
        rt = spectrum.getRT()
        spectrum_as_dict = {
            "index":index,
            "mz"  :mzs,
            "i"    :Is,
            "rt"   :rt,
            "level":spectrum.getMSLevel()
        }
        list_of_spectra.append(spectrum_as_dict)
        index += 1
        
    return list_of_spectra

def Convert_MSExperiment_To_Dict(msexperiment_obj, lower_index):
    print("Start Converting from MSExperiment object to python dictionary...")

    pyopenms.MzMLFile().store(f"temp_{lower_index}.mzML" ,msexperiment_obj)

    lines = []
    with open(f"temp_{lower_index}.mzML", 'r') as infile:
        for line in infile:
            lines.append(line)

    msexperiment_dict = {
        "from": lower_index,
        "to": lower_index + msexperiment_obj.getNrSpectra(),
        "mzml": lines,
        "level":msexperiment_obj.getSpectrum(0).getMSLevel()
    }

    os.remove(f"temp_{lower_index}.mzML")

    return msexperiment_dict

def Write_Json(python_object):
    print("Start writing the spectra to json...")

    file_name = f'level_{python_object[0]["level"]}.json'
    with open("meta_inf.txt", 'a') as outfile:
        outfile.write(f'Level:{python_object[0]["level"]}\n')
        for object in python_object:
            outfile.write(f'{object["from"]},{object["to"]}\n')
    with open(file_name, 'w') as outfile:
        json.dump(python_object, outfile)
    
    print("A file has been written successfully... ")
    return "A file has been written successfully... "


def Write_Spectra_To_Json(msexperiment_obj, number_of_slaves):
    print("Start the process of converting mzml to json...")

    number_of_spectra_each_file_hold = msexperiment_obj.getNrSpectra() // number_of_slaves
    lower_indices = [i*number_of_spectra_each_file_hold for i in range(number_of_slaves)]
    msexperiment_objects = []
    for i in range(number_of_slaves):
        exp = pyopenms.MSExperiment()
        for j in range(i*number_of_spectra_each_file_hold, (i+1) * number_of_spectra_each_file_hold):
            spectrum = msexperiment_obj.getSpectrum(j)
            exp.addSpectrum(spectrum)

        msexperiment_objects.append(exp)

    with concurrent.futures.ThreadPoolExecutor() as exe:
      results = exe.map(Convert_MSExperiment_To_Dict, msexperiment_objects, lower_indices)

    msexperiment_dicts = [result for result in results] 
    Write_Json(msexperiment_dicts)


def Mzml_To_Json(mzml_file, number_of_slaves=4):
    print("Start reading the file...")

    experiment_obj = pyopenms.MSExperiment()
    pyopenms.MzMLFile().load(mzml_file, experiment_obj)
    level_1_experiment_obj = pyopenms.MSExperiment()
    level_2_experiment_obj = pyopenms.MSExperiment()
    for spectrum in experiment_obj.getSpectra():
        if spectrum.getMSLevel() == 1:
            level_1_experiment_obj.addSpectrum(spectrum)
        else:
            level_2_experiment_obj.addSpectrum(spectrum)
    number_of_slaves_to_hold_level_1 = number_of_slaves 
    number_of_slaves_to_hold_level_2 = number_of_slaves 

    number_of_slaves_to_hold_levels = [number_of_slaves_to_hold_level_1, number_of_slaves_to_hold_level_2]
    experiment_objects = [level_1_experiment_obj, level_2_experiment_obj]

    source_files = experiment_obj.getSourceFiles()[0]
    file_name = source_files.getNameOfFile()[0:source_files.getNameOfFile().find('.')]
    file_id = source_files.getNativeIDTypeAccession()

    with open("meta_inf.txt", 'w') as outfile:
        outfile.write(f'{file_name}\n{file_id}\n{experiment_obj.getNrSpectra()}\n')

    for i in range(len(experiment_objects)):
        Write_Spectra_To_Json(experiment_objects[i], number_of_slaves_to_hold_levels[i])



def Get_Search_Results(mzml_file, fasta_file, sample_id="sample"):
    obj = PeptideIdentification(mzml_file, fasta_file)

    obj.Set_Protease('Trypsin')
    obj.Set_Missed_Cleavages(2)
    obj.Set_Modifications(fixed=[b'Carbamidomethyl (C)'], variable=[b'Oxidation (M)', b'Acetyl (N-term)'])
    obj.Set_Mass_Tolerance(20, 4.5)
    obj.Set_Charges(2, 6)
    obj.Search(sample_id)

    return obj.df

def Peptide_Search(mzml_json, fasta_file):
    ## Convert from json to mzml...
    with open(mzml_json) as infile:
        mzml_data = json.load(infile)
    
    with open(f'searchfile_{mzml_data["From"]}.mzML', 'w') as outfile:
        for line in mzml_data["MZML"]:
            outfile.write(line)
    
    results_df = Get_Search_Results(f'searchfile_{mzml_data["From"]}.mzML', fasta_file)

    results_df.to_csv(f'results_{mzml_data["From"]}.csv') 

    os.remove(f'searchfile_{mzml_data["From"]}.mzML')
   

parser = argparse.ArgumentParser()
parser.add_argument("--fasta_id", type=str)
parser.add_argument("--output_dir", type=str)
parser.add_argument("--mzml_to_json", type=str)
parser.add_argument("--peptideSearch", type=str)
args = parser.parse_args()

if args.fasta_id:
    if args.output_dir:
        Get_Protein_Fasta_By_Id(args.fasta_id, args.output_dir)
    else:
        Get_Protein_Fasta_By_Id(args.fasta_id)
    
if args.mzml_to_json:
    Mzml_To_Json(args.mzml_to_json)


if args.peptideSearch:
    files = args.peptideSearch.split(',')
    Peptide_Search(files[0], files[1])





 
    
    


