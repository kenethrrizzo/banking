package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

	logger.Debug("Customers returned.")
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

	logger.Debug("Customer found.")
	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:root@/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
