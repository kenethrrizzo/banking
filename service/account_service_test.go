package service

import (
	"testing"

	"github.com/kenethrrizzo/banking/dto"
)

func Test_should_return_a_validation_error_response_when_the_request_is_not_validated(t *testing.T) {
	request := dto.NewAccountRequest{
		CustomerId: "100",
		Type:       "saving",
		Amount:     0,
	}
	service := NewAccountService(nil)

	_, err := service.NewAccount(request)

	if err == nil {
		t.Error("Failed while testing the new account validation")
	}
}
