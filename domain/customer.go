package domain

// Secundary port: Interface

import (
	"github.com/kenethrrizzo/banking/dto"
	errs "github.com/kenethrrizzo/banking/error"
)

type Customer struct {
	Id          string `db:"Id"`
	Name        string `db:"Name"`
	City        string `db:"City"`
	ZipCode     string `db:"ZipCode"`
	DateOfBirth string `db:"DateOfBirth"`
	Status      string `db:"Status"`
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}

func (c Customer) statusAsText() string {
	if c.Status == "N" {
		return "Inactive"
	} else {
		return "Active"
	}
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}
