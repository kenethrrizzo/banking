package main

import (
	"log"
	"net/http"
)

func Start() {
	// routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// listen
	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}