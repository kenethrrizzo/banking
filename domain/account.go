package domain

import (
	"github.com/kenethrrizzo/banking/dto"
	errs "github.com/kenethrrizzo/banking/error"
)

type Account struct {
	Id          string  `db:"Id"`
	CustomerId  string  `db:"CustomerId"`
	OpeningDate string  `db:"OpeningDate"`
	Type        string  `db:"Type"`
	Amount      float64 `db:"Amount"`
	Status      string  `db:"Status"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{Id: a.Id}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
