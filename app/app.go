package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	hndl "github.com/kenethrrizzo/banking/app/handlers"
	mdl "github.com/kenethrrizzo/banking/app/middlewares"
	repo "github.com/kenethrrizzo/banking/domain/repositories"
	"github.com/kenethrrizzo/banking/logger"
	"github.com/kenethrrizzo/banking/service"
)

func Start() {
	router := mux.NewRouter()
	serverConfig := NewServerConfig()
	dbclient := getDatabaseClient()

	customerepodb := repo.NewCustomerRepositoryDb(dbclient)
	accountrepodb := repo.NewAccountRepositoryDb(dbclient)

	cushandl := hndl.CustomerHandler{
		Service: service.NewCustomerService(customerepodb),
	}

	acchandl := hndl.AccountHandler{
		Service: service.NewAccountService(accountrepodb),
	}

	// Routes
	router.
		HandleFunc("/customers", cushandl.GetAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")

	router.
		HandleFunc("/customers/{customer-id:[0-9]+}", cushandl.GetCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")

	router.
		HandleFunc("/customers/{customer-id:[0-9]+}/account", acchandl.NewAccount).
		Methods(http.MethodPost).
		Name("NewAccount")

	router.
		HandleFunc("/customers/{customer-id:[0-9]+}/account/{account-id:[0-9]+}", acchandl.MakeTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	// Middleware
	am := mdl.AuthMiddleware{
		Repo: repo.NewAuthRepository(),
	}
	router.Use(am.AuthorizationHandler())

	// Listen
	logger.Error(http.ListenAndServe(
		fmt.Sprintf("%s:%s",
			serverConfig.Address,
			serverConfig.Port),
		router).Error())
}

func getDatabaseClient() *sqlx.DB {
	dbconfig := NewDatabaseConfig()

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
