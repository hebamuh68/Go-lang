package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}


		//========================================================= call slave of grade 1
		if (string(body)=="g1"){
			fmt.Printf("Data Requested: Grade 1 \n")
			resp, err := http.Get("http://127.0.0.1:8088")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, "%s", string(body))

			
		//========================================================= call slave of grade 2	
		} else if (string(body)=="g2"){
			fmt.Printf("Data Requested: Grade 2\n")
			resp, err := http.Get("http://127.0.0.1:8099")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, "%s", string(body))
			
		}	
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}

