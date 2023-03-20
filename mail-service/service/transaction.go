package service

import (
	"github.com/danisbagus/go-messaging-kafka/mail-service/model"
	"github.com/danisbagus/go-messaging-kafka/mail-service/repo"
)

type ITransactionService interface {
	SendTransactionPaid(data model.Transaction) error
}

type TransactionService struct {
	TransactionRepo repo.ITransactionRepo
}

func NewTransactionService(transactionRepo repo.ITransactionRepo) ITransactionService {
	return &TransactionService{
		TransactionRepo: transactionRepo,
	}
}

func (s TransactionService) SendTransactionPaid(data model.Transaction) error {
	return s.TransactionRepo.SendTransactionPaid(data)
}
