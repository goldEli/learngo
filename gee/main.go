package gee

import (
	"log"
	"net/http"
)

type (
	HandlerFunc func(*Context)
	Engine      struct {
		*RouterGroup
		router *Router
		groups []*RouterGroup
	}
)

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.router.AddRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.router.AddRoute("POST", pattern, handler)
}

func (engine *Engine) Run(port string) {
	log.Fatal(http.ListenAndServe(port, engine))
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	c := newContext(w, r)

	engine.router.Handle(c)
}
