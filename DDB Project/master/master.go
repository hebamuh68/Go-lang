package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"sql.go/web/Methods"
)

func main() {
	// The handler function that take requests and parameters from the clients...
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		//=============================  Parse the request body as JSON
		var requestBody Methods.RequestBody
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		Feature := requestBody.Feature

		//=============================  Parse the request body as JSON
		switch Feature {
		case "Search_fasta":
			fastaID := requestBody.FastaID
			metaFasta := Methods.Get_Fasta_Id(fastaID)

			if metaFasta.Fasta_id == "" {
				// The following code responsible for downloading the unexisting fasta file, store it in a slave, then send it to 
				// the client...
				Methods.Get_New_Fasta(fastaID) // fetch the fasta file from uniprot...
				metaFasta.Fasta_id = fastaID
				files, _ := os.ReadDir(Methods.Data_dir) // list the files of the directory which has the working files...
				for _, file := range files {
					if strings.Contains(file.Name(), ".fasta") {
						file_path := Methods.Data_dir + "/" + file.Name() // get the fasta file path...
						_, name, header, sequence := Methods.Read_Fasta(file_path) // parse the file path...
						metaFasta.Name = name
						metaFasta.Slave = Methods.The_Next_Slave_Fasta()
						
						metaFasta.Insert_Fasta_Id() // Insert the meta data related to that file into master database...

						// Send the fasta data to the determined slave...
						Methods.Send_Fasta_To_Slave(metaFasta.Slave, metaFasta.Fasta_id, name, header, sequence)

						err := os.Remove(file_path) // remove the fasta file from master OS...
						if err != nil {
							panic(err.Error())
						}

						Methods.Send_Fasta(metaFasta.Slave, fastaID, w) // Send the fasta file to the client...
						return
					}
				}
			}
			// if the file does exist in one of the slaves then send it directly to the client...
			Methods.Send_Fasta(metaFasta.Slave, fastaID, w)

		case "Mzml":
			msSampleID := requestBody.MsSampleID
			fastaID := requestBody.FastaID
			metaFasta := Methods.Get_Fasta_Id(fastaID)

			if metaFasta.Fasta_id == "" {
				fmt.Println("Not exist")
				return
			}

			fastaIDBytes := []byte(metaFasta.Fasta_id)

			//============================= Make HTTP GET request to another service
			
			// Get the wanted fasta file from the hosting slave then route it to other slaves to start the peptide search process...
			resp := Methods.Get_Fasta_From_Slave(metaFasta.Slave, fastaIDBytes)

			defer resp.Body.Close()
			//=========================================================================================================
			//============================= Send fasta to other slaves

			// Recieving and routing the fasta file...
			len := resp.ContentLength
			body := make([]byte, len)
			resp.Body.Read(body)

			var fastaFile Methods.Fasta

			// Converting the json message into Fasta object
			json.Unmarshal(body, &fastaFile)
			if err != nil {
				http.Error(w, "Failed to parse request body", http.StatusBadRequest)
				return
			}

			// Get the level2 meta data to map the parameters to each slave to start the peptide identification process...
			sampleMetaData := Methods.Get_Level_2(msSampleID)
			var search_results [4][]Methods.Record

			for _, object := range sampleMetaData {
				// Start the mapping...
				go Methods.Map_Data_To_Slave(object.Slave, msSampleID, object.From_spectra, fastaFile, w)
			}

			// Recieving the results from the slaves...
			search_results[0] = <-Methods.Slave_1
			search_results[1] = <-Methods.Slave_2
			search_results[2] = <-Methods.Slave_3
			search_results[3] = <-Methods.Slave_4

			var final_results []Methods.Record

			// Reducing the outputs from the mapping process...
			for i := range search_results {
				for j, _ := range search_results[i] {
					final_results = append(final_results, search_results[i][j])
				}
			}


			Methods.Write_CSV(final_results, "SearchRes.csv")
			// Load the CSV file contents into a variable
			csvData, err := ioutil.ReadFile("SearchRes.csv")
			if err != nil {
				http.Error(w, "Failed to read CSV file", http.StatusInternalServerError)
				return
			}

			// Write the CSV data to the response writer
			w.Header().Set("Content-Type", "text/csv")
			w.Header().Set("Content-Disposition", "attachment; filename=SearchRes.csv")
			w.Write(csvData)

			fmt.Println("Done...")
		}
	})

	fmt.Println("Master starting on port 8080")
	http.ListenAndServe("localhost:8080", nil)
}
