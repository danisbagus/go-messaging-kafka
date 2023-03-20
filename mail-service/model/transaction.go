package model

type Transaction struct {
	TransactionID string   `json:"transaction_id"`
	Customer      Customer `json:"customer"`
}

type TransactionPayload struct {
	Action string      `json:"action"`
	Data   Transaction `json:"data"`
}

type Customer struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
