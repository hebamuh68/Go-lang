package main

import (
	"html/template"
	"net/http"
   // "fmt"
)

type prodSpec struct {
	Size   string
	Weight float32
	Descr  string
}

type product struct {
	ProdID int
	// Name   string
	Name  string
	Cost  float64
	Specs prodSpec
}


var tpl *template.Template
var prod1 product
var name = "Heba"

func main(){

    prod1 = product{
		ProdID: 15,
		Name:   "Wicked Cool Phone",
		Cost:   899,
		Specs: prodSpec{
			Size:   "150 x 70 x 7 mm",
			Weight: 65,
			Descr:  "Over priced shiny thing designed to shatter on impact",
		},
	}

	tpl,_ = tpl.ParseGlob("templates/*.html")
    http.HandleFunc("/welcome",welcomeHandler)
    http.HandleFunc("/prod",prodHandler)
    http.HandleFunc("/bye",byeHandler)

    http.ListenAndServe(":8080", nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index1.html", name)
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index2.html", name)
}

func prodHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index3.html", prod1)
}