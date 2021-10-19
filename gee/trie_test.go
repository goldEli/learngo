package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParsePattern(t *testing.T) {
	type item struct {
		pattern string
		parts   []string
	}
	data := []item{
		{"/p/:name", []string{"p", ":name"}},
		{"/p/*", []string{"p", "*"}},
		{"/p/*name/123", []string{"p", "*name"}},
	}

	for _, d := range data {
		ok := reflect.DeepEqual(parsePattern(d.pattern), d.parts)
		if !ok {
			t.Fatal(fmt.Printf("test parsePattern failed: %v != %v", parsePattern(d.pattern), d.parts))
		}
	}
}

func TestGetRoute(t *testing.T) {
	type item struct {
		method  string
		pattern string
		param   map[string]string
		url     string
	}

	data := []item{
		{url: "/", method: "GET", pattern: "/", param: map[string]string{}},
		{url: "/hello/geektutu", method: "GET", pattern: "/hello/:name", param: map[string]string{"name": "geektutu"}},
		{url: "/hello/b/c", method: "GET", pattern: "/hello/b/c", param: map[string]string{}},
		{url: "/assets/geektutu", method: "GET", pattern: "/assets/*filepath", param: map[string]string{"filepath": "geektutu"}},
	}

	r := newRouter()
	for _, d := range data {
		r.AddRoute(d.method, d.pattern, nil)
	}

	for _, d := range data {
		n, ps := r.getRoute(d.method, d.url)

		if n == nil {
			t.Fatal(fmt.Printf("nil shouldn't be returned: %v", d.pattern))
		}

		if n.pattern != d.pattern {
			t.Fatal(fmt.Printf("should match %v: %v", d.pattern, n.pattern))
		}

		ok := reflect.DeepEqual(d.param, ps)
		if !ok {
			t.Fatal(fmt.Printf("%v, test GetRoute failed: %v != %v", d.pattern, d.param, ps))
		}
	}

}
