package service

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/danisbagus/go-messaging-kafka/transaction-service/dto"
	"github.com/danisbagus/go-messaging-kafka/transaction-service/model"
	"github.com/danisbagus/go-messaging-kafka/transaction-service/repo"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type ITransactionService interface {
	Create(*dto.NewTransactionRequest) (*dto.NewTransactionResponse, error)
}

type TransactionService struct {
	MailRepo repo.IMailRepo
}

func NewTransactionService(mailRepo repo.IMailRepo) ITransactionService {
	return &TransactionService{
		MailRepo: mailRepo,
	}
}

func (r TransactionService) Create(form *dto.NewTransactionRequest) (*dto.NewTransactionResponse, error) {
	transactionID := fmt.Sprintf("TR%v", String(6))
	customerName := "danisbagus22"
	customerEmail := "danishter22@gmail.com"

	transaction := model.Transaction{
		TransactionID:   transactionID,
		ProductID:       form.ProductID,
		Quantity:        form.Quantity,
		TransactionDate: time.Now(),
		Customer: model.Customer{
			Email: customerEmail,
			Name:  customerName,
		},
	}

	transactionPayload := model.TransactionPayload{
		Action: "create",
		Data:   transaction,
	}

	mailData := model.Mail{
		Topic: os.Getenv("TOPIC_TRANSACTION"),
		Key:   fmt.Sprintf("%s-%s", customerName, transactionID),
		Data:  transactionPayload,
	}

	go r.MailRepo.SendMail(&mailData)

	response := dto.NewNewTransactionResponse(&transaction)

	return response, nil
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
