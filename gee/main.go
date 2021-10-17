package gee

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)
type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRouter("GET", pattern, handler)
}

func (engine *Engine) Run(port string) {
	log.Fatal(http.ListenAndServe(port, engine))
}

func (engine *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "_" + pattern
	engine.router[key] = handler
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	pattern := r.URL.Path
	key := method + "_" + pattern

	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %q:%q\n", method, pattern)
	}
}
