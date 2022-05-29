package service

import (
	"time"

	"github.com/kenethrrizzo/banking/domain"
	"github.com/kenethrrizzo/banking/dto"
	errs "github.com/kenethrrizzo/banking/error"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	account := domain.Account{
		Id:          "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format(time.RFC3339),
		Type:        req.Type,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}
