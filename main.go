package main

import (
	"fmt"
	"net/http"

	"github.com/goldEli/learngo/gee"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "/: hello")
	})
	engine.Run(":9999")
}
