package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)
type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.router.AddRouter("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.router.AddRouter("POST", pattern, handler)
}

func (engine *Engine) Run(port string) {
	log.Fatal(http.ListenAndServe(port, engine))
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	c := newContext(w, r)

	engine.router.Handle(c)
}
