package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"sql.go/web/Methods"
)

func main() {
	// Handles incoming requests to the root path ("/storeFasta") which is responsible for storing the fasta file into the DataBase...
	http.HandleFunc("/storeFasta", func(w http.ResponseWriter, r *http.Request) {
		// Read the request body and convert it into bytes...
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)

		var fasta_data Methods.Fasta

		// Converting the json message into Fasta object
		json.Unmarshal(body, &fasta_data)

		// Insert the new Fasta object into the slave database
		fasta_data.Insert_Fasta_Data()

		fmt.Println("Done Storing fasta with id..." + fasta_data.Fasta_id)
	})
	// Handles incoming requests to the root path ("/") which is responsible for routing the fasta file to the client...
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		fasta_id := string(body)
		fasta_data := Methods.Get_Fasta_Data(fasta_id) // Get the data from the DataBase...

		// Write the fasta file to the HTTP response writer...
		json_value, _ := json.Marshal(&fasta_data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json_value)

	})

	// Handles incoming requests to the root path ("/peptideSearch") which is responsible
	//for runing the peptide search process and send the results to the reducer...
	http.HandleFunc("/peptideSearch", func(w http.ResponseWriter, r *http.Request) {
		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		//=============================  Parse the request body as JSON
		// The request body contains the mzml sample id , fasta file, and other DataBase query parameters...
		var requestBody Methods.RequestBody
		_ = json.Unmarshal(body, &requestBody)

		msSampleID := requestBody.MsSampleID
		index := requestBody.Index
		fastaFile := requestBody.FastaFile
		// The file paths which are the parameters to the python tool...
		fasta_file_path := "slave_2_search.fasta"
		json_file_path := "slave_2_searchfile.json"

		mzml_object := Methods.Get_Spectra(msSampleID, index) // Get the mzml data from the DataBase...

		// Creating the fasta and json files...
		Methods.Write_Fasta(fastaFile.Header, fastaFile.Sequence, Methods.Data_dir+"/"+fasta_file_path)
		Methods.Write_Json(mzml_object, Methods.Data_dir+"/"+json_file_path)

		tool_dir := Methods.Data_dir
		command_str := fmt.Sprintf("python proteomics_pipeline.py --peptideSearch %s,%s", json_file_path, fasta_file_path)
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute) 
		defer cancel()
		
		cmd := exec.CommandContext(ctx, "bash", "-c", command_str)
		cmd.Dir = tool_dir
		err = cmd.Run()
		if err != nil {
			panic(err.Error())
		}

		os.Remove(Methods.Data_dir + "/" + fasta_file_path)
		os.Remove(Methods.Data_dir + "/" + json_file_path)

		files, _ := os.ReadDir(Methods.Data_dir)

		// Converting the results into json values to send them back to the reducer...
		for _, file := range files {
			ind := fmt.Sprint(index)
			if strings.Contains(file.Name(), ".csv") && strings.Contains(file.Name(), "_"+ind) {
				csv_file_path := Methods.Data_dir + "/" + file.Name()

				csv_record := Methods.CSVtoJSON(csv_file_path)

				os.Remove(csv_file_path)

				jsonData, err := json.Marshal(csv_record)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				// Writing the results back...
				w.Header().Set("Content-Type", "application/json")
				fmt.Println("Send searh results data...")
				w.Write(jsonData)

				break
			}
		}

	})

	// Handles incoming requests to the root path ("/storeSpectra") which is responsible for storing sending sample from master
	// into the slave DataBase...
	http.HandleFunc("/storeSample", func(w http.ResponseWriter, r *http.Request) {
		// Read the request body
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)

		var sample_data Methods.MsSample

		// Converting the json message into Fasta object
		json.Unmarshal(body, &sample_data)

		// Insert the new Fasta object into the slave database
		sample_data.Insert_SSample()

		fmt.Println("Done Storing Sample with id..." + sample_data.Sample_id)
	})

	// Handles incoming requests to the root path ("/storeSpectra") which is responsible for storing sending spectra from master
	// into the slave DataBase...
	http.HandleFunc("/storeSpectra/", func(w http.ResponseWriter, r *http.Request) {
		id := path.Base(r.URL.Path) // read the id parameter which is sent into the url...

		// Reading the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		var mzml_object4 Methods.MzML
		var spectrum4 Methods.SpectrumData

		// Converting the json message into MzML object
		json.Unmarshal(body, &mzml_object4)
		// fmt.Println(id, mzml_object4.From)

		// Storing the spectrum object into DataBase...
		spectrum4.Sample_id = id
		spectrum4.Spectrum_index = mzml_object4.From
		spectrum4.Spectrum = mzml_object4
		spectrum4.Insert_Spectra()

		fmt.Println("Done Storing Spectra...")
		r.Body.Close()

	})

	// Listen on port 8088
	fmt.Println("Slave_2 starting on port 8092")
	if err := http.ListenAndServe("127.0.0.1:8092", nil); err != nil {
		log.Fatal(err)
	}
}
