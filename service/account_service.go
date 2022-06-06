package service

import (
	"fmt"
	"strings"
	"time"

	repo "github.com/kenethrrizzo/banking/domain/repositories"
	ent "github.com/kenethrrizzo/banking/domain/entities"
	"github.com/kenethrrizzo/banking/dto"
	errs "github.com/kenethrrizzo/banking/error"
	"github.com/kenethrrizzo/banking/logger"
)

const (
	WITHDRAWAL string = "withdrawal"
	DEPOSIT    string = "deposit"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo repo.AccountRepository
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
	account := ent.Account{
		Id:          "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format(dbTSLayout),
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

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	t := ent.Transaction{
		AccountId: req.AccountId,
		Amount:    req.Amount,
		Type:      req.Type,
		Date:      time.Now().Format(dbTSLayout),
	}

	err := t.ValidateType()
	if err != nil {
		logger.Error(fmt.Sprintf("Transaction type '%s' is not allowed", req.Type))
		return nil, errs.NewValidationError("Transaction type not allowed")
	}

	err = t.ValidateAmount()
	if err != nil {
		logger.Error("Amount must be positive")
		return nil, errs.NewValidationError("Negative amount")
	}

	if t.IsWithdrawal() {
		acc, err := s.repo.FindById(t.AccountId)
		if err != nil {
			return nil, err
		}
		if !acc.CanWithdraw(t.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	transaction, apperr := s.repo.SaveTransaction(t)
	if apperr != nil {
		return nil, apperr
	}

	response := transaction.ToDto()
	logger.Debug(fmt.Sprint("TransactionResponse: ", response))
	return &response, nil
}

func NewAccountService(repo repo.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
