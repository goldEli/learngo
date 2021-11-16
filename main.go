package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/goldEli/learngo/geecache"
)

var db = map[string]string{
	"tom":  "43",
	"lily": "53",
	"sam":  "63",
}

func main() {
	key := "scores"
	geecache.NewGroup(key, 2<<10, geecache.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))
	addr := "localhost:9999"
	peers := geecache.NewHTTPPool(addr)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))

}
