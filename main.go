package main

import (
	"log"
	"net/http"

	"github.com/goldEli/learngo/gee"
)

func main() {
	engine := new(gee.Engine)

	log.Fatal(http.ListenAndServe(":9999", engine))
}
