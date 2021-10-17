package gee

import (
	"fmt"
	"net/http"
)

type Engine struct {
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "url = %q\n", r.URL.Path)
	default:
		fmt.Fprintf(w, "404 not fount: %q\n", r.URL.Path)
	}
}
