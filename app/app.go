package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// routes
	router.HandleFunc(
		"/greet", greet).Methods(http.MethodGet)
	router.HandleFunc(
		"/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc(
		"/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc(
		"/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// listen
	log.Fatalln(http.ListenAndServe("localhost:8080", router))
}
