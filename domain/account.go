package domain

import errs "github.com/kenethrrizzo/banking/error"

type Account struct {
	Id          string
	CustomerId  string
	OpeningDate string
	Type        string
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
