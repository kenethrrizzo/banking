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

func (a Account) CanWithdraw(amount float64) bool {
	return amount < a.Amount
}

type AccountRepository interface {
	FindById(string) (*Account, *errs.AppError)
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}
