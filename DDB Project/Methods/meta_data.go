package Methods

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	var err error
	Meta_db, err = sql.Open("mysql", "heba:heba333@tcp(localhost:3306)/MetaData")
	if err != nil {
		fmt.Println("Not Connected...")
		panic(err.Error())
	}

	fmt.Println("Connected...")
}

func Get_Fasta_Id(id string) (metafasta MetaFasta) {
	query_statement := "SELECT * FROM MetaData.Fasta WHERE Fasta_id = ?"
	result, err := Meta_db.Query(query_statement, id)
	if err != nil {
		fmt.Println("Not Executed...")
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&metafasta.Fasta_id, &metafasta.Name, &metafasta.Slave)
		if err != nil {
			panic(err.Error())
		}
	}
	return
}

func (metafasta *MetaFasta) Insert_Fasta_Id() {
	query_statement := "INSERT INTO MetaData.Fasta VALUES(?, ?, ?);"
	_, err := Meta_db.Exec(query_statement, metafasta.Fasta_id, metafasta.Name, metafasta.Slave)
	if err != nil {
		fmt.Println("Not Inserted...")
		panic(err.Error())
	}

}

func (metafasta *MetaFasta) Update_Fasta_Id(id string) {
	query_statement := "UPDATE MetaData.Fasta SET Fasta_id = ?, Name = ?, Slave = ? WHERE Fasta_id = ?"
	_, err := Meta_db.Exec(query_statement, metafasta.Fasta_id, metafasta.Name, metafasta.Slave, id)
	if err != nil {
		fmt.Println("Not Updated...")
		panic(err.Error())
	}
}

func (metafasta *MetaFasta) Delete_Fasta_Id() {
	query_statement := "DELETE FROM MetaData.Fasta WHERE Fasta_id = ?"
	_, err := Meta_db.Exec(query_statement, metafasta.Fasta_id)
	if err != nil {
		fmt.Println("Not Deleted...")
		panic(err.Error())
	}
}

func Get_Sample(id string) (ms_sample MsSample) {
	query_statement := "SELECT * FROM MetaData.MsSample WHERE Sample_id = ?"
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

func (ms_sample *MsSample) Insert_Sample() {
	query_statement := "INSERT INTO MetaData.MsSample(Sample_id, Name, NumberOfSpectra) VALUES(?, ?, ?);"
	_, err := Meta_db.Exec(query_statement, ms_sample.Sample_id, ms_sample.Name, ms_sample.Number_of_spectra)
	if err != nil {
		fmt.Println("Not Inserted...")
		panic(err.Error())
	}
}

func (ms_sample *MsSample) Delete_Sample() {
	query_statement := "DELETE FROM MetaData.MsSample WHERE Sample_id = ?"
	_, err := Meta_db.Exec(query_statement, ms_sample.Sample_id)
	if err != nil {
		fmt.Println("Not Deleted...")
		panic(err.Error())
	}
}

func (ms_sample *MsSample) Update_Sample(id string) {
	query_statement := "UPDATE MetaData.MsSample SET Sample_id = ?, Name = ?, NumberOfSpectra = ? WHERE Sample_id = ?"
	_, err := Meta_db.Exec(query_statement, ms_sample.Sample_id, ms_sample.Name, ms_sample.Number_of_spectra, id)
	if err != nil {
		fmt.Println("Not Updated...")
		panic(err.Error())
	}
}

func Get_Level_1(sample_id string) (level_1 []Level1) {
	query_statement := "SELECT Level1_id, Slave, FromSpectra, ToSpectra FROM Level1 WHERE Sample_id = ?"
	result, err := Meta_db.Query(query_statement, sample_id)
	if err != nil {
		fmt.Println("Not Executed...")
		panic(err.Error())
	}
	defer result.Close()
	index := 0
	for result.Next() {
		var lev_1 Level1
		err := result.Scan(&lev_1.Level1_id, &lev_1.Slave, &lev_1.From_spectra, &lev_1.To_spectra)
		if err != nil {
			panic(err.Error())
		}
		level_1 = append(level_1, lev_1)
		index++
	}
	return
}

func (level_1 *Level1) Insert_Level_1(sample_id string) {
	query_statement := "INSERT INTO Level1(Slave, FromSpectra, ToSpectra, Sample_id) VALUES(?, ?, ?, ?);"
	_, err := Meta_db.Exec(query_statement, level_1.Slave, level_1.From_spectra, level_1.To_spectra, sample_id)
	if err != nil {
		fmt.Println("Not Inserted...")
		panic(err.Error())
	}
}

func (level_1 *Level1) Update_Level_1(sample_id string) {
	query_statement := "UPDATE Level1 SET Level1_id = ? Sample_id = ?, Slave = ?, FromSpectra = ?, ToSpectra = ? WHERE Sample_id = ?"
	_, err := Meta_db.Exec(query_statement, level_1.Level1_id, sample_id, level_1.Slave, level_1.From_spectra, level_1.To_spectra, sample_id)
	if err != nil {
		fmt.Println("Not Updated...")
		panic(err.Error())
	}
}

func (Level1 *Level1) Delete_Level_1() {
	query_statement := "DELETE FROM Level1 WHERE Level1_id = ?"
	_, err := Meta_db.Exec(query_statement, Level1.Level1_id)
	if err != nil {
		fmt.Println("Not Deleted...")
		panic(err.Error())
	}
}

func Get_Level_2(sample_id string) (level_2 []Level2) {
	query_statement := "SELECT Level2_id, Slave, FromSpectra, ToSpectra FROM MetaData.Level2 WHERE Sample_id = ?"
	result, err := Meta_db.Query(query_statement, sample_id)
	if err != nil {
		fmt.Println("Not Executed...")
		panic(err.Error())
	}
	defer result.Close()
	index := 0
	for result.Next() {
		var lev_2 Level2
		err := result.Scan(&lev_2.Level2_id, &lev_2.Slave, &lev_2.From_spectra, &lev_2.To_spectra)
		if err != nil {
			panic(err.Error())
		}
		level_2 = append(level_2, lev_2)
		index++
	}
	return
}

func (level_2 *Level2) Insert_Level_2(sample_id string) {
	query_statement := "INSERT INTO MetaData.Level2(Slave, FromSpectra, ToSpectra, Sample_id) VALUES(?, ?, ?, ?);"
	_, err := Meta_db.Exec(query_statement, level_2.Slave, level_2.From_spectra, level_2.To_spectra, sample_id)
	if err != nil {
		fmt.Println("Not Inserted...")
		panic(err.Error())
	}
}

func (level_2 *Level2) Update_Level_2(sample_id string) {
	query_statement := "UPDATE MetaData.Level2 SET Level2_id = ? Sample_id = ?, Slave = ?, FromSpectra = ?, ToSpectra = ? WHERE Sample_id = ?"
	_, err := Meta_db.Exec(query_statement, level_2.Level2_id, sample_id, level_2.Slave, level_2.From_spectra, level_2.To_spectra, sample_id)
	if err != nil {
		fmt.Println("Not Updated...")
		panic(err.Error())
	}
}

func (level_2 *Level2) Delete_Level_2() {
	query_statement := "DELETE FROM MetaData.Level2 WHERE Level2_id = ?"
	_, err := Meta_db.Exec(query_statement, level_2.Level2_id)
	if err != nil {
		fmt.Println("Not Deleted...")
		panic(err.Error())
	}
}

