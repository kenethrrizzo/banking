package dto

type TransactionRequest struct {
	CustomerId string  `json:"customer_id"`
	AccountId  string  `json:"account_id"`
	Type       string  `json:"type"`
	Amount     float64 `json:"amount"`
}
