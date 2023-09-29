package request_model

type PaymentTransaction struct {
	TransactionId string `json:"transaction_id"`
	Amount        int    `json:"amount"`
}
