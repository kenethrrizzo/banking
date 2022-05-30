package dto

type TransactionResponse struct {
	Id        string  `json:"transaction_id"`
	Type      string  `json:"transaction_type"`
	Date      string  `json:"transaction_date"`
	AccountId string  `json:"account_id"`
	Amount    float64 `json:"account_new_balance"`
}
