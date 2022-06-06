package entities

import (
	"errors"

	"github.com/kenethrrizzo/banking/dto"
)

const (
	WITHDRAWAL string = "withdrawal"
	DEPOSIT    string = "deposit"
)

type Transaction struct {
	Id        string  `db:"Id"`
	AccountId string  `db:"AccountId"`
	Amount    float64 `db:"Amount"`
	Type      string  `db:"Type"`
	Date      string  `db:"Date"`
}

func (t Transaction) IsWithdrawal() bool {
	return t.Type == WITHDRAWAL
}

func (t Transaction) IsDeposit() bool {
	return t.Type == DEPOSIT
}

func (t Transaction) IsAmountPositive() bool {
	return t.Amount > 0
}

func (t Transaction) ValidateType() error {
	if !t.IsDeposit() && !t.IsWithdrawal() {
		return errors.New("Invalid transaction type")
	}
	return nil
}

func (t Transaction) ValidateAmount() error {
	if !t.IsAmountPositive() {
		return errors.New("Invalid amount")
	}
	return nil
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		Id:        t.Id,
		Type:      t.Type,
		Date:      t.Date,
		AccountId: t.AccountId,
		Amount:    t.Amount,
	}
}
