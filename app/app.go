package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenethrrizzo/banking/config"
	"github.com/kenethrrizzo/banking/domain"
	"github.com/kenethrrizzo/banking/logger"
	"github.com/kenethrrizzo/banking/service"
	"github.com/spf13/viper"
)

func Start() {
	router := mux.NewRouter()
	config.InitializeDatabaseConfig()

	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	dbDomain := viper.GetString("database.domain")
	dbPort := viper.GetString("database.port")

	logger.Error(http.ListenAndServe(fmt.Sprintf("%s:%s", dbDomain, dbPort), router).Error())
}
