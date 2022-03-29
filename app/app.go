package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenethrrizzo/banking/domain"
	"github.com/kenethrrizzo/banking/service"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatalln(http.ListenAndServe("localhost:8080", router))
}
