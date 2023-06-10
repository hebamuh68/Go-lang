package Methods

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Read_Fasta(file_path string) (id string, name string, description string, sequence string) {

	handle, err := os.Open(file_path)

	if err != nil {
		panic(err)
	}
	defer handle.Close()

	content := bufio.NewScanner(handle)
	content.Scan()
	description = content.Text()
	id_start_index := strings.Index(description, " RecName")
	id = description[1:id_start_index]
	name_start_index := strings.Index(description, "=")
	name_end_index := strings.Index(description, "{")
	if name_end_index == -1 {
		name_end_index = strings.Index(description, ";")
	}

	name = description[name_start_index+1 : name_end_index]
	for content.Scan() {
		sequence += content.Text()
	}

	return id, name, description, sequence
}

func Write_Fasta(description string, sequence string, file_path string) {
	handle, err := os.Create(file_path)
	if err != nil {
		panic(err)
	}

	defer handle.Close()
	_, err2 := handle.WriteString(description + "\n" + sequence)
	if err2 != nil {
		panic(err2)
	}

}

func Write_CSV(search_results []Record, file_path string) {
	csvFile, err := os.Create(file_path)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	row := []string{"Sample", "Spectrum_index", "Precursor_mass", "Intensity_sum", "Intensity_mean", "RT", "Matched_peptide"}
	_ = csvwriter.Write(row)
	for _, record := range search_results {
		Sample := record.Sample
		Spectrum_index := fmt.Sprintf("%v", record.Spectrum_index)
		Precursor_mass := fmt.Sprintf("%v", record.Precursor_mass)
		Intensity_sum := fmt.Sprintf("%v", record.Intensity_sum)
		Intensity_mean := fmt.Sprintf("%v", record.Intensity_mean)
		RT := fmt.Sprintf("%v", record.RT)
		Matched_peptide := record.Matched_peptide

		row := []string{Sample, Spectrum_index, Precursor_mass, Intensity_sum, Intensity_mean, RT, Matched_peptide}
		_ = csvwriter.Write(row)
	}
	defer csvwriter.Flush()
}

func Read_Json(file_path string) (mzml []MzML) {
	json_file, _ := os.Open(file_path)
	defer json_file.Close()
	json_data, _ := ioutil.ReadAll(json_file)
	json.Unmarshal(json_data, &mzml)

	return mzml
}
