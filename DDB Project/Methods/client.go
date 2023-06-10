package Methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Master_Connect(id string) { //192.168.43.140
	master_url := fmt.Sprintf("http://127.0.0.1:8080/getfasta/%s", id)
	_, err := http.Get(master_url)
	if err != nil {
		fmt.Println("Failed to connect...")
		panic(err.Error())
	}
}

func Client_Get_Spectra(sample_id string, level, from_spectra, to_spectra int) {
	master_url := fmt.Sprintf("http://127.0.0.1:8080/getspectra/?id=%s&level=%d&from=%d&to=%d", sample_id, level, from_spectra, to_spectra)

	_, err := http.Get(master_url)
	if err != nil {
		fmt.Println("Failed to connect...")
		panic(err.Error())
	}
}

func Handle_Post(conn *gin.Context) {
	var fasta_data Fasta
	if err := conn.BindJSON(&fasta_data); err != nil {
		return
	}
	Write_Fasta(fasta_data.Header, fasta_data.Sequence, "prot.fasta")
	conn.IndentedJSON(http.StatusCreated, fasta_data)
}

func Handle_Spectra(conn *gin.Context) {
	var spectra []Spectrum
	if err := conn.BindJSON(&spectra); err != nil {
		return
	}
	output, err := json.Marshal(&spectra)
	if err != nil {
		panic(err.Error())
	}
	err = ioutil.WriteFile("spectra.json", output, 0644)
	if err != nil {
		panic(err.Error())
	}
	conn.IndentedJSON(http.StatusCreated, spectra)
}


//MAIN FUNCTIONS ====================================================================================================================
//====================================================================================================================

func Search_fasta_Handler(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {

			
		//============================= Read the selected value from the form data
			feature := "Search_fasta"
			fastaId := r.FormValue("fastaId")

		//============================= Create the request body based on the selected value
			MasterRequestBody := RequestBody{
								Feature:      feature,
								FastaID:         fastaId,
							}
			requestBody, err := json.Marshal(MasterRequestBody)
			if err != nil {
				panic(err)
			}

		//============================= Make the HTTP request to the master
			req, err := http.NewRequest("POST", "http://127.0.0.1:8080", bytes.NewBuffer(requestBody))
			if err != nil {
				panic(err)
			}
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()


		//============================= Read the response body
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}


		//============================= Parse the HTML template and replace {{.}} with the response body
			tmpl := template.Must(template.ParseFiles("templates/Search_fasta.html"))
			err = tmpl.Execute(w, string(respBody))
			if err != nil {
				panic(err)
			}
			
		} else {


		//============================= Render the HTML form on GET request
			tmpl := template.Must(template.ParseFiles("templates/Search_fasta.html"))
			err := tmpl.Execute(w, nil)
			if err != nil {
				panic(err)
			}
		}
}


func Mzml_Handler(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {

			
		//============================= Read the selected value from the form data
			msSampleId := r.FormValue("msSampleId")
			fastaId := r.FormValue("fastaId")
			feature := "Mzml"

			//fmt.Println(fastaId)

		//============================= Create the request body based on the selected value
			requestBody := []byte(fmt.Sprintf("{\"msSampleId\": \"%s\", \"fastaId\": \"%s\",\"feature\": \"%s\"}", msSampleId, fastaId, feature))


		//============================= Make the HTTP request to the master
			req, err := http.NewRequest("POST", "http://127.0.0.1:8080", bytes.NewBuffer(requestBody))
			if err != nil {
				panic(err)
			}
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()


		//============================= Read the response body
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}


		//============================= Parse the HTML template and replace {{.}} with the response body
			tmpl := template.Must(template.ParseFiles("templates/Mzml.html"))
			err = tmpl.Execute(w, string(respBody))
			if err != nil {
				panic(err)
			}
			
		} else {


		//============================= Render the HTML form on GET request
			tmpl := template.Must(template.ParseFiles("templates/Mzml.html"))
			err := tmpl.Execute(w, nil)
			if err != nil {
				panic(err)
			}
		}
}