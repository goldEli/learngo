package gee

import (
	"fmt"
	"net/http"
)

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}

func (r *Router) AddRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "_" + pattern
	r.handlers[key] = handler
}

func (r *Router) Handle(c *Context) {
	key := c.Method + "_" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.Status(http.StatusNotFound)
		fmt.Fprintf(c.Writer, "404 NOT FOUND: %q:%q\n", c.Method, c.Path)
	}
}
