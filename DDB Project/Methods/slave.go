package Methods

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Send_Fasta_To_Client(conn *gin.Context) {
	var meta_data FastaMetaData // master msg
	var fasta_data Fasta

	fasta_data = Get_Fasta_Data(meta_data.Fasta_id) //
	json_value, _ := json.Marshal(&fasta_data)
	_, err := http.Post("http://192.168.1.4:8008:8008/fasta", "application/json", bytes.NewBuffer(json_value))
	if err != nil {
		panic(err.Error())
	}
}

func Store_And_Send_Fasta_To_Client(conn *gin.Context) {
	var fasta_data Fasta
	if err := conn.BindJSON(&fasta_data); err != nil {
		return
	}

	fasta_data.Insert_Fasta_Data() //imp

	json_value, _ := json.Marshal(&fasta_data)
	_, err := http.Post("http://192.168.1.4:8008/fasta", "application/json", bytes.NewBuffer(json_value))
	if err != nil {
		panic(err.Error())
	}

}

func Store_Spectra(conn *gin.Context) {
	sample_id := conn.Param("sample_id")
	var mzml_obj MzML

	if err := conn.BindJSON(&mzml_obj); err != nil {
		return
	}
	var spec SpectrumData
	spec.Sample_id = sample_id
	spec.Spectrum_index = mzml_obj.From
	spec.Spectrum = mzml_obj

	spec.Insert_Spectra()
	fmt.Println("Done haha...")
}

func Store_Sample(conn *gin.Context) {
	var ms_sample MsSample
	if err := conn.BindJSON(&ms_sample); err != nil {
		return
	}

	ms_sample.Insert_Sample()
}

func Handle_Send_Spectra(conn *gin.Context) {
	sample_id := conn.Param("id")
	var meta_data SpectraMetaData
	var spectra []MzML
	if err := conn.BindJSON(&meta_data); err != nil {
		return
	}

	for i := meta_data.From; i <= meta_data.To; i++ {
		spec := Get_Spectra(sample_id, i)
		spectra = append(spectra, spec.Spectrum)
	}

	json_value, _ := json.Marshal(&spectra)
	_, err := http.Post("http://192.168.1.4:8008:8008/spectra", "application/json", bytes.NewBuffer(json_value))
	if err != nil {
		panic(err.Error())
	}

}

func Write_Json(spectrum_data SpectrumData, file_path string) {
	specJSON, err := json.Marshal(spectrum_data.Spectrum)
	if err != nil {
		panic(err.Error())
	}

	_ = ioutil.WriteFile(file_path, specJSON, 0644)

}

func CSVtoJSON(csvPath string) []Record {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var rec Record
	var records []Record

	for _, each := range csvData[1:] {
		rec.Sample = each[1]
		rec.Spectrum_index, _ = strconv.Atoi(each[2])
		rec.Precursor_mass, _ = strconv.ParseFloat(each[3], 64)
		rec.Intensity_sum, _ = strconv.ParseFloat(each[4], 64)
		rec.Intensity_mean, _ = strconv.ParseFloat(each[5], 64)
		rec.RT, _ = strconv.ParseFloat(each[6], 64)
		rec.Matched_peptide = each[7]
		records = append(records, rec)
	}

	return records
	// jsonData, err := json.Marshal(records)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// return jsonData

	// Write JSON data to file
	// file, err := os.Create("outputData.json")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	// _, err = file.Write(jsonData)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// } else {
	// 	fmt.Println("json file created successfully!")
	// }

}
