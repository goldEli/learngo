package gee

import (
	"fmt"
	"net/http"
	"strings"
)

// roots key eg, roots['GET'] roots['POST]
// handlers key eg, handlers['GET-/p/:lang/doc'] handlers['POST-/p/book']

type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)

	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (r *Router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	// handle params
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
			}
		}
	}
	return n, params
}

func (r *Router) AddRouter(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "_" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *Router) Handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "_" + n.pattern
		r.handlers[key](c)
	} else {
		c.Status(http.StatusNotFound)
		fmt.Fprintf(c.Writer, "404 NOT FOUND: %q:%q\n", c.Method, c.Path)
	}

}
