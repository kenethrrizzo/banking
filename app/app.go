package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenethrrizzo/banking/config"
	"github.com/kenethrrizzo/banking/domain"
	"github.com/kenethrrizzo/banking/logger"
	"github.com/kenethrrizzo/banking/service"
)

func Start() {
	router := mux.NewRouter()

	dbconfig := config.NewDatabaseConfig()

	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	logger.Error(http.ListenAndServe(fmt.Sprintf("%s:%s", dbconfig.Domain, dbconfig.Port), router).Error())
}
