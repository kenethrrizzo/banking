package domain

import errs "github.com/kenethrrizzo/banking/error"

type Customer struct {
	Id          string `db:"cu_id"`
	Name        string `db:"cu_name"`
	City        string `db:"cu_city"`
	ZipCode     string `db:"cu_zipcode"`
	DateOfBirth string `db:"cu_date_of_birth"`
	Status      string `db:"cu_status"`
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
