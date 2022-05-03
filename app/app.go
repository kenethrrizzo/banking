package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenethrrizzo/banking/domain"
	"github.com/kenethrrizzo/banking/logger"
	"github.com/kenethrrizzo/banking/service"
)

func Start() {
	router := mux.NewRouter()

	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	logger.Error(http.ListenAndServe("localhost:8080", router).Error())
}
