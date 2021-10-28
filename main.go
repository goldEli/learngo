package main

import (
	"net/http"

	"github.com/goldEli/learngo/gee"
)

func main() {
	r := gee.New()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
