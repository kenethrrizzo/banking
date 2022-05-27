package domain

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kenethrrizzo/banking/config"
	errs "github.com/kenethrrizzo/banking/error"
	"github.com/kenethrrizzo/banking/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	var findAllQuery string
	if status == "" {
		findAllQuery = "select cu_id, cu_name, cu_city, cu_zipcode, cu_date_of_birth, cu_status from customers"
	} else {
		findAllQuery = "select cu_id, cu_name, cu_city, cu_zipcode, cu_date_of_birth, cu_status from customers where cu_status = '" + status + "'"
	}
	err := d.client.Select(&customers, findAllQuery)

	if err != nil {
		logger.Error("Error while querying customer table -> " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	logger.Info("Customers found.")
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	var customer Customer
	findByIdQuery := "select cu_id, cu_name, cu_city, cu_zipcode, cu_date_of_birth, cu_status from customers where cu_id = ?"

	err := d.client.Get(&customer, findByIdQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Customer not found -> " + err.Error())
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customer -> " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	logger.Info("Customer found.")
	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbconfig := config.NewDatabaseConfig()

	client, err := sqlx.Open(dbconfig.Driver, fmt.Sprintf("%s:%s@/%s", dbconfig.Username, dbconfig.Password, dbconfig.Name))
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
