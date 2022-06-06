package domain

import "testing"

func TestShouldReturnErrorWhenTransactionTypeIsNotDepositOrWithdrawl(t *testing.T) {
	// Arrange
	transaction := Transaction{
		Type: "InvalidTransactionType",
	}
	// Act
	err := transaction.ValidateType()
	// Assert
	if err == nil {
		t.Error("Method doesn't return error")
	}
	if err.Error() != "Invalid transaction type" {
		t.Error("Invalid message while testing transaction type")
	}
}

func TestShouldReturnErrorWhenAmountIsLessThanZero(t *testing.T) {
	transaction := Transaction{
		Amount: -20,
	}
	err := transaction.ValidateAmount()
	if err == nil {
		t.Error("Method doesn't return error")
	}
	if err.Error() != "Invalid amount" {
		t.Error("Invalid message while testing amount")
	}
}