package dto

import (
	"time"

	"github.com/danisbagus/go-messaging-kafka/transaction-service/model"
)

type NewTransactionResponse struct {
	TransactionID   string    `json:"transaction_id"`
	ProductID       int64     `json:"product_id"`
	Quantity        int64     `json:"quantity"`
	TransactionDate time.Time `json:"transaction_date"`
}

type NewTransactionRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

func NewNewTransactionResponse(data *model.Transaction) *NewTransactionResponse {
	result := &NewTransactionResponse{
		TransactionID:   data.TransactionID,
		ProductID:       data.ProductID,
		Quantity:        data.Quantity,
		TransactionDate: data.TransactionDate,
	}

	return result
}
