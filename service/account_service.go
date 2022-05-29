package service

import (
	"strings"
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
	//validations
	if req.Amount < 5000 {
		return nil, errs.NewValidationError("To open a new account, you need to deposit atleast 5000.00")
	}
	if strings.ToLower(req.Type) != "saving" && strings.ToLower(req.Type) != "checking" {
		return nil, errs.NewValidationError("Account type should be 'checking' or 'saving'")
	}
	//map dto to account
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

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
