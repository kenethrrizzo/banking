package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/kenethrrizzo/banking/config"
	"github.com/kenethrrizzo/banking/domain"
	"github.com/kenethrrizzo/banking/logger"
	"github.com/kenethrrizzo/banking/service"
)

func Start() {
	router := mux.NewRouter()
	dbconfig := config.NewDatabaseConfig()
	dbclient := getDatabaseClient()

	customerepodb := domain.NewCustomerRepositoryDb(dbclient)
	//accountrepodb := domain.NewCustomerRepositoryDb(dbclient)

	cushandl := CustomerHandler{service.NewCustomerService(customerepodb)}

	router.HandleFunc("/customers", cushandl.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", cushandl.getCustomer).Methods(http.MethodGet)

	logger.Error(http.ListenAndServe(fmt.Sprintf("%s:%s", dbconfig.Domain, dbconfig.Port), router).Error())
}

func getDatabaseClient() *sqlx.DB {
	dbconfig := config.NewDatabaseConfig()

	client, err := sqlx.Open(dbconfig.Driver, fmt.Sprintf("%s:%s@/%s", dbconfig.Username, dbconfig.Password, dbconfig.Name))
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}