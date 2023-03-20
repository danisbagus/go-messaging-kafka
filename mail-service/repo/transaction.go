package repo

import (
	"github.com/danisbagus/go-messaging-kafka/mail-service/model"
	"github.com/danisbagus/go-messaging-kafka/mail-service/modules"
)

type ITransactionRepo interface {
	SendTransactionPaid(data model.Transaction) error
}

type TransactionRepo struct {
	smtp modules.Smtp
}

func NewTransactionRepo(smpt modules.Smtp) *TransactionRepo {
	return &TransactionRepo{
		smtp: smpt,
	}
}

func (r TransactionRepo) SendTransactionPaid(data model.Transaction) error {
	option := modules.SmtpOption{
		Subject:      "Transaction Paid",
		Target:       []string{data.Customer.Email},
		TemplateData: data,
		FileNames:    "mail-service/templates/transaction_paid.html",
	}
	if err := r.smtp.SendMail(option); err != nil {
		return err
	}

	return nil
}
