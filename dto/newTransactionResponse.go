package dto

type NewTransactionResponse struct {
	TransactionID string  `json:"transaction_id"`
	AcctAmount    float64 `json:"balance"`
}
