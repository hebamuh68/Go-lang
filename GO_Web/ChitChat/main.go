package main

import (
    "net/http"
)

func main(){

    // The multiplexer
    mux := http.NewServeMux()
    mux.HandleFunc("/", index)

    // Serving static files
    files := http.FileServer(http.Dir("/public"))
    mux.Handle("/static/", http.StripPrefix("/static/", files))



}