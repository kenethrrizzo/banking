package domain

import (
	"github.com/kenethrrizzo/banking/dto"
	errs "github.com/kenethrrizzo/banking/error"
)

type Account struct {
	Id          string
	CustomerId  string
	OpeningDate string
	Type        string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{Id: a.Id}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
