package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {

			
		//============================= Read the selected value from the form data
			selectedValue := r.FormValue("requestBody")


		//============================= Create the request body based on the selected value
			requestBody := []byte(selectedValue)


		//============================= Make the HTTP request to the server
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
			tmpl := template.Must(template.ParseFiles("client1.html"))
			err = tmpl.Execute(w, string(respBody))
			if err != nil {
				panic(err)
			}
			
		} else {


		//============================= Render the HTML form on GET request
			tmpl := template.Must(template.ParseFiles("client1.html"))
			err := tmpl.Execute(w, nil)
			if err != nil {
				panic(err)
			}
		}
	})

	fmt.Println("Client_1 starting on port 8091")
	http.ListenAndServe("127.0.0.1:8091", nil)
}
