package main

import "fmt"

type Response struct {
	ID   string
	Type string
}

type SpecificResponse struct {
	Response // Response type is embedded inside SpecificResponse
	Data     string
}

func main() {
	sr := SpecificResponse{Data: "123", Response: Response{ID: "123"}}
	fmt.Printf(sr.ID)
	fmt.Printf(sr.Response.ID)
}
