package model

import "time"

type Transaction struct {
	TransactionID   string    `json:"transaction_id"`
	ProductID       int64     `json:"product_id"`
	Quantity        int64     `json:"quantity"`
	TransactionDate time.Time `json:"transaction_date"`
	Customer        Customer  `json:"customer"`
}

type TransactionPayload struct {
	Action string      `json:"action"`
	Data   Transaction `json:"data"`
}

type Customer struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
