package main
import(
    "fmt"
    "net/http"
)

/*
This function handles the incoming requests and sends the response. 
It takes two parameters: a ResponseWriter interface that can be used 
to send an HTTP response to the client, and a Request struct that 
represents the client's HTTP request.
*/
func handler(writer http.ResponseWriter, request * http.Request) {
    fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1: ])
}
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8090", nil)
}