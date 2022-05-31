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
	serverConfig := config.NewServerConfig()
	dbclient := getDatabaseClient()

	customerepodb := domain.NewCustomerRepositoryDb(dbclient)
	accountrepodb := domain.NewAccountRepositoryDb(dbclient)

	cushandl := CustomerHandler{
		service.NewCustomerService(customerepodb),
	}

	acchandl := AccountHandler{
		service.NewAccountService(accountrepodb),
	}

	router.HandleFunc(
		"/customers", cushandl.getAllCustomers,
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/customers/{customer-id:[0-9]+}", cushandl.getCustomer,
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/customers/{customer-id:[0-9]+}/account", acchandl.newAccount,
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/customers/{customer-id:[0-9]+}/account/{account-id:[0-9]+}", acchandl.makeTransaction,
	).Methods(http.MethodPost)

	logger.Error(http.ListenAndServe(
		fmt.Sprintf("%s:%s",
			serverConfig.Address,
			serverConfig.Port),
		router).Error())
}

func getDatabaseClient() *sqlx.DB {
	dbconfig := config.NewDatabaseConfig()

	client, err := sqlx.Open(
		dbconfig.Driver,
		fmt.Sprintf("%s:%s@/%s",
			dbconfig.Username,
			dbconfig.Password,
			dbconfig.Name))

	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
