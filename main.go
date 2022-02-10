package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {
	// routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// listen
	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}

func greet(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello, World!")
}

func getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Keneth", City: "Guayaquil", Zipcode: "112345"},
		{Name: "Camila", City: "Durán", Zipcode: "112335"},
	}

	rw.Header().Add("Content-Type", "application/json")

	json.NewEncoder(rw).Encode(customers)
}
