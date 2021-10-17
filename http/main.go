package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "Hello, %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}

}

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// })
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	for k, v := range r.Header {
	// 		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	// 	}
	// })
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
