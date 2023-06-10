import pandas
import pyopenms
from AnalyzeSpectrum import AnalyzeSpectrum
import numpy

class PeptideIdentification(object):
    def __init__(self, mzml_file, target_protein_fasta_file):
        self.search_engine = pyopenms.SimpleSearchEngineAlgorithm()
        self.search_engine_parameters = self.search_engine.getDefaults()
        self.protein_ids = []
        self.matched_peptide_ids = []
        self.mzml_file = mzml_file
        self.target_protein_fasta_file = target_protein_fasta_file
        self.exp = pyopenms.MSExperiment()
        pyopenms.MzMLFile().load(self.mzml_file, self.exp)
        self.df = pandas.DataFrame(columns=['Sample', 'Spectrum_index', 'Precursor_mass', 'Intensity_sum', 'Intensity_mean',
                                            'RT', 'Matched_peptide'])
        
        
    def Set_Protease(self, protease='Trypsin'):
        self.search_engine_parameters[b'enzyme'] = protease
        self.search_engine.setParameters(self.search_engine_parameters)
    
    def Set_Mass_Tolerance(self, precursor_mass_tolerance=10, fragment_mass_tolerance=10):
        self.search_engine_parameters[b'precursor:mass_tolerance'] = float(precursor_mass_tolerance)
        self.search_engine_parameters[b'fragment:mass_tolerance'] = float(fragment_mass_tolerance)
        self.search_engine.setParameters(self.search_engine_parameters)
    
    def Set_Charges(self, min_charge=2, max_charge=5):
        self.search_engine_parameters[b'precursor:min_charge'] = min_charge
        self.search_engine_parameters[b'precursor:max_charge'] = max_charge
        self.search_engine.setParameters(self.search_engine_parameters)
        
    def Set_Isotopes(self, isotopes=[0, 1]):
        self.search_engine_parameters[b'precursor:isotopes'] = isotopes
        self.search_engine.setParameters(self.search_engine_parameters)
    
    def Set_Peptide_Sizes(self, min_size=7, max_size=40):
        self.search_engine_parameters[b'peptide:min_size'] = min_size
        self.search_engine_parameters[b'peptide:max_size'] = max_size
        self.search_engine.setParameters(self.search_engine_parameters)
    
    def Set_Missed_Cleavages(self, missed_cleavages_number=1):
        self.search_engine_parameters[b'peptide:missed_cleavages'] = missed_cleavages_number
        self.search_engine.setParameters(self.search_engine_parameters)
    
    def Set_Modifications(self, fixed=[b'Carbamidomethyl (C)'], variable=[b'Oxidation (M)']):
        self.search_engine_parameters[b'modifications:fixed'] = fixed
        self.search_engine_parameters[b'modifications:variable'] = variable
        self.search_engine.setParameters(self.search_engine_parameters)
    
    def Search(self, sample_number="1"):
        self.search_engine.search(self.mzml_file, self.target_protein_fasta_file,
                                  self.protein_ids, self.matched_peptide_ids)
        
        samples = [sample_number] * len(self.matched_peptide_ids)
        pre_masses = []
        spectrum_indices = []
        i_summes = []
        i_means = []
        rts = []
        peptide_seqs = []
        
        for peptide_identified in self.matched_peptide_ids:
            print(f"Sample{sample_number}:", end=' ')
            print(f"Observed Spectrum at index {peptide_identified.getMetaValue('scan_index')} has M/Z =",end=' ')
            print(f"{peptide_identified.getMZ()} and {len(peptide_identified.getHits())} hits")
            
            spectrum = self.exp.getSpectrum(peptide_identified.getMetaValue('scan_index'))
            i_summes.append(numpy.sum(spectrum.get_peaks()[1]))
            i_means.append(numpy.mean(spectrum.get_peaks()[1]))
            pre_masses.append(peptide_identified.getMZ())
            spectrum_indices.append(peptide_identified.getMetaValue('scan_index'))
            rts.append(peptide_identified.getRT())
            
            for hit in peptide_identified.getHits():
                peptide_seqs.append(str(hit.getSequence()))
        
        data = {'Sample':samples, 'Spectrum_index':spectrum_indices, 'Precursor_mass':pre_masses, 'Intensity_sum':i_summes,
                'Intensity_mean':i_means, 'RT':rts, 'Matched_peptide':peptide_seqs}
        self.df = pandas.DataFrame(data)
        
        
        
    
    def Generate_Theoretical_Spectrum(self, peptide_seq, min_charge=1, max_charge=1):
        tsg = pyopenms.TheoreticalSpectrumGenerator()
        theoretical_spectrum = pyopenms.MSSpectrum()
        parameters = pyopenms.Param()
        
        parameters.setValue("add_metainfo", "true")
        tsg.setParameters(parameters)
        peptide = pyopenms.AASequence.fromString(peptide)
        
        tsg.getSpectrum(theoretical_spectrum, peptide, min_charge, max_charge)
        
        return theoretical_spectrum
        
    def Show_Peptides_Distribution(self, msexperiment_object=None, given_matched_peptides=None, folder_path=None, sample='0'):
            if given_matched_peptides == None:
                if len(self.matched_peptide_ids) == 0:
                    print("Matched Peptides Not Found...")
                    return
                else:
                    spectrums_indices = sorted(self.df['Spectrum_index'])

                    obj = AnalyzeSpectrum()
                    mz, i, rt = obj.Get_MSData(self.exp, spectrums_indices)
                    obj.Draw_3D_Spectrum(mz, i, rt, sample=sample, folder_path=folder_path)
            else:
                if len(given_matched_peptides) == 0:
                    print("Matched Peptides Not Found...")
                    return
                else:
                    spectrums_indices = sorted([peptide_identified.getMetaValue('scan_index') for peptide_identified in given_matched_peptides])

                    obj = AnalyzeSpectrum()
                    mz, i, rt = obj.Get_MSData(msexperiment_object, spectrums_indices)
                    obj.Draw_3D_Spectrum(mz, i, rt)
    