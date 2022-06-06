package service

// Primary port: Service

import (
	repo "github.com/kenethrrizzo/banking/domain/repositories"
	"github.com/kenethrrizzo/banking/dto"
	errs "github.com/kenethrrizzo/banking/error"
)

//go:generate mockgen -destination=../mocks/service/mock_customer_service.go -package=service github.com/kenethrrizzo/banking/service CustomerService
type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

// Service implementation (core - business logic)

type DefaultCustomerService struct {
	repo repo.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, customer := range customers {
		response = append(response, customer.ToDto())
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	response := customer.ToDto()
	return &response, nil
}
func NewCustomerService(repo repo.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
