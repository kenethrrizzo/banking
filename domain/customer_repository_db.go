package domain

// Adapter: Database

import (
	"database/sql"
	"fmt"

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
		findAllQuery = "select Id, Name, City, ZipCode, DateOfBirth, Status from Customers"
	} else {
		findAllQuery = fmt.Sprintf("select Id, Name, City, ZipCode, DateOfBirth, Status from Customers where Status = '%s'", status)
	}
	err := d.client.Select(&customers, findAllQuery)

	if err != nil {
		logger.Error("Error while querying customer table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	logger.Info("Customers found.")
	logger.Debug(fmt.Sprint("Customers:", customers))
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	var customer Customer
	findByIdQuery := "select Id, Name, City, ZipCode, DateOfBirth, Status from Customers where Id = ?"

	err := d.client.Get(&customer, findByIdQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Customer not found: " + err.Error())
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customer: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	logger.Info("Customer found.")
	logger.Debug(fmt.Sprint("Customer:", customer))
	return &customer, nil
}

func NewCustomerRepositoryDb(dbclient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbclient}
}
