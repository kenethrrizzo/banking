package domain

import "testing"

func TestShouldReturnErrorWhenTransactionTypeIsNotDepositOrWithdrawl(t *testing.T) {
	// Arrange
	transaction := Transaction{
		Type: "InvalidTransactionType",
	}
	// Act
	apperr := transaction.ValidateType()
	// Assert
	if apperr.Error() != "Invalid transaction type" {
		t.Error("Invalid message while testing transaction type")
	}
}