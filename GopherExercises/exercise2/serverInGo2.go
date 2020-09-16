package main

// ServeMux is an HTTP request multiplexer.
// It matches the URL of each incoming request against a list of
// registered patterns and calls the handler for the pattern that most closely matches the URL.
// ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil,
// which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:
// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	//r.URL.Path replies with the URL path in your request.
}
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love boo and string %s!", r.URL.Path[1:])
	//r.URL.Path replies with the URL path in your request.
}
func main() {
	http.HandleFunc("/", handler) //the http.HandleFunc, puts in the params and calls the handler func.
	http.HandleFunc("/boo", handler1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// req.FormValue
//passing function reference inside a function.
