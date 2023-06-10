package main

import (
	//"fmt"
	//"os"
	//"os/exec"

	//"fmt"

	"sql.go/web/Methods"
)

func Populate_Spectra_Data(){
	meta_inf_path := Methods.Data_dir + "/" + "meta_inf.txt"
	Methods.Populate_Spectra_Meta_Data(meta_inf_path)

	sample := Methods.Get_Sample("MS:1000777")
	Methods.Post_Ms_Sample_V1(sample)

	Methods.Post_Spectra_To_Slaves("MS:1000777")
}

func Populate_Fasta_Data(file_path string){
	id, name, header, seq := Methods.Read_Fasta(file_path)

	var meta_fasta Methods.MetaFasta

	meta_fasta.Fasta_id = id
	meta_fasta.Name = name
	meta_fasta.Slave = Methods.The_Next_Slave_Fasta()

	meta_fasta.Insert_Fasta_Id()

	Methods.Send_Fasta_To_Slave(meta_fasta.Slave, id, name, header, seq)
}

func main() {

	bsa_file_path := Methods.Data_dir + "/" + "bsa.fasta"
	soat_file_path := Methods.Data_dir + "/" + "SOAT1_HUMAN.fasta"

	Populate_Fasta_Data(bsa_file_path)
	Populate_Fasta_Data(soat_file_path)

	Populate_Spectra_Data()
	
	//fmt.Print(Methods.Get_Fasta_Data("P02769"))

}
