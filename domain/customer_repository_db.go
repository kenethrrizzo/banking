package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	errs "github.com/kenethrrizzo/banking/error"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var findAllQuery string

	if status != "" {
		if status == "V" || status == "N" {
			findAllQuery = "select cu_id, cu_name, cu_city, cu_zipcode, cu_date_of_birth, cu_status from customers where cu_status = '" + status + "'"
		} else {
			log.Println("Error -> Values not recognized")
			return nil, errs.NewUnexpectedError("Unexpected values in args")
		}
	} else {
		findAllQuery = "select cu_id, cu_name, cu_city, cu_zipcode, cu_date_of_birth, cu_status from customers"
	}
	

	rows, err := d.client.Query(findAllQuery)
	if err != nil {
		log.Println("Error while querying customer table ->", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers ->", err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	findByIdQuery := "select cu_id, cu_name, cu_city, cu_zipcode, cu_date_of_birth, cu_status from customers where cu_id = ?"

	row := d.client.QueryRow(findByIdQuery, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		log.Println("Error while scanning customer ->", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:root@/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
