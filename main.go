package main

import (
	"github.com/goldEli/learngo/gee"
)

func main() {
	r := gee.New()
	///Users/miaoyu/Desktop/test/learngo/static

	r.Static("/assets", "./static")

	r.Run(":9999")
}
