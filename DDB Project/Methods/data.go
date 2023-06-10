package Methods

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	var err error
	Data_db, err = sql.Open("mysql", "heba:heba333@tcp(localhost:3306)/ProteinData")
	if err != nil {
		fmt.Println("Not Connected...")
		panic(err.Error())
	}
	fmt.Println("Connected...")
}

func Get_Fasta_Data(id string) (fasta Fasta) {
	query_statement := "SELECT * FROM ProteinData.Fasta WHERE Fasta_id = ?"
	result, err := Data_db.Query(query_statement, id)
	if err != nil {
		fmt.Println("Not Executed...")
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&fasta.Fasta_id, &fasta.Name, &fasta.Header, &fasta.Sequence)
		if err != nil {
			panic(err.Error())
		}
	}
	return
}

func (fasta *Fasta) Insert_Fasta_Data() {
	query_statement := "INSERT INTO ProteinData.Fasta VALUES(?, ?, ?, ?);"
	_, err := Data_db.Exec(query_statement, fasta.Fasta_id, fasta.Name, fasta.Header, fasta.Sequence)
	if err != nil {
		fmt.Println("Not Inserted...")
		panic(err.Error())
	}

}

func (fasta *Fasta) Update_Fasta_Data(id string) {
	query_statement := "UPDATE ProteinData.Fasta SET Fasta_id = ?, Name = ?, Header = ?, Sequence = ? WHERE Fasta_id = ?"
	_, err := Data_db.Exec(query_statement, fasta.Fasta_id, fasta.Name, fasta.Header, fasta.Sequence, id)
	if err != nil {
		fmt.Println("Not Updated...")
		panic(err.Error())
	}
}

func (fasta *Fasta) Delete_Fasta_Data() {
	query_statement := "DELETE FROM ProteinData.Fasta WHERE Fasta_id = ?"
	_, err := Data_db.Exec(query_statement, fasta.Fasta_id)
	if err != nil {
		fmt.Println("Not Deleted...")
		panic(err.Error())
	}
}

func Get_SSample(id string) (ms_sample MsSample) {
	query_statement := "SELECT * FROM ProteinData.MsSample WHERE Sample_id = ?"
	result, err := Data_db.Query(query_statement, id)
	if err != nil {
		fmt.Println("Not Executed...")
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&ms_sample.Sample_id, &ms_sample.Name, &ms_sample.Number_of_spectra)
		if err != nil {
			panic(err.Error())
		}
	}
	return
}

func (ms_sample *MsSample) Insert_SSample() {
	query_statement := "INSERT INTO ProteinData.MsSample(Sample_id, Name, NumberOfSpectra) VALUES(?, ?, ?);"
	_, err := Meta_db.Exec(query_statement, ms_sample.Sample_id, ms_sample.Name, ms_sample.Number_of_spectra)
	if err != nil {
		fmt.Println("Not Inserted...")
		panic(err.Error())
	}
}

func (ms_sample *MsSample) Delete_SSample() {
	query_statement := "DELETE FROM ProteinData.MsSample WHERE Sample_id = ?"
	_, err := Meta_db.Exec(query_statement, ms_sample.Sample_id)
	if err != nil {
		fmt.Println("Not Deleted...")
		panic(err.Error())
	}
}

func (ms_sample *MsSample) Update_SSample(id string) {
	query_statement := "UPDATE ProteinData.MsSample SET Sample_id = ?, Name = ?, NumberOfSpectra = ? WHERE Sample_id = ?"
	_, err := Meta_db.Exec(query_statement, ms_sample.Sample_id, ms_sample.Name, ms_sample.Number_of_spectra, id)
	if err != nil {
		fmt.Println("Not Updated...")
		panic(err.Error())
	}
}

func Get_Spectra(sample_id string, index int) (spectra SpectrumData) {
	rows, err := Data_db.Query("SELECT * FROM ProteinData.Spectra WHERE Sample_id = ? AND SpectrumIndex = ?", sample_id, index)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var spect string
		err := rows.Scan(&spectra.Spectrum_index, &spect, &spectra.Sample_id)
		if err != nil {
			panic(err.Error())
		}
		json.Unmarshal([]byte(spect), &spectra.Spectrum)
	}
	return
}

func (spectra *SpectrumData) Insert_Spectra() {
	query_statement := "INSERT INTO ProteinData.Spectra(SpectrumIndex, MZML, Sample_id) VALUES(?, ?, ?);"

	specJSON, err := json.Marshal(spectra.Spectrum)
	if err != nil {
		panic(err.Error())
	}
	_, err = Data_db.Exec(query_statement, spectra.Spectrum_index, string(specJSON), spectra.Sample_id)
	if err != nil {
		fmt.Println("Not Inserted...")
		panic(err.Error())
	}
}

func (spectra *SpectrumData) Delete_Spectra() {
	query_statement := "DELETE FROM ProteinData.Spectra WHERE Sample_id = ? AND SpectrumIndex = ?"
	_, err := Data_db.Exec(query_statement, spectra.Sample_id, spectra.Spectrum_index)
	if err != nil {
		fmt.Println("Not Deleted...")
		panic(err.Error())
	}
}

func (spectra *SpectrumData) Update_Spectra(sample_id string, spectrum_index int) {
	query_statement := "UPDATE ProteinData.Spectra SET Sample_id = ?, SpectrumIndex = ?, MZML = ? WHERE Sample_id = ? AND SpectrumIndex = ?"
	_, err := Data_db.Exec(query_statement, spectra.Sample_id, spectra.Spectrum_index, spectra.Spectrum, sample_id, spectrum_index)
	if err != nil {
		fmt.Println("Not Updated...")
		panic(err.Error())
	}
}
