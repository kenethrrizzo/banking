package dto

type NewAccountRequest struct {
	CustomerId string  `json:"customer_id"`
	Type       string  `json:"type"`
	Amount     float64 `json:"amount"`
}
