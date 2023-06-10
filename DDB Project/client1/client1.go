package main

import (
	"fmt"
	"html/template"
	"net/http"

	"sql.go/web/Methods"
)

var tpl *template.Template

func main() {
	
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/", main_Handler)
    http.HandleFunc("/Search_fasta", Methods.Search_fasta_Handler)
    http.HandleFunc("/Mzml", Methods.Mzml_Handler)
    http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))

	fmt.Println("Client_1 starting on port 8090")
	http.ListenAndServe("127.0.0.1:8090", nil)
}


func main_Handler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

