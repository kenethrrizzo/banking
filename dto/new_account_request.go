package dto

type NewAccountRequest struct {
	CustomerId string  `json:"cu_id"`
	Type       string  `json:"ac_type"`
	Amount     float64 `json:"ac_amount"`
}
