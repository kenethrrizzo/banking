package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"FullName"`
	City    string `json:"city" xml:"City"`
	Zipcode string `json:"zip_code" xml:"ZipCode"`
}

func getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Keneth", City: "Guayaquil", Zipcode: "112345"},
		{Name: "Camila", City: "Durán", Zipcode: "112335"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)
	} else {
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(customers)
	}
}

func greet(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello, World!")
}