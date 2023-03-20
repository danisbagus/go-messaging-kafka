package repo

import (
	"encoding/json"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/go-messaging-kafka/transaction-service/model"
	"github.com/danisbagus/go-messaging-kafka/transaction-service/modules"
	"go.uber.org/zap"
)

type IMailRepo interface {
	SendMail(*model.Mail)
}

type MailRepo struct {
	kp *modules.KafkaProducer
}

func NewMailRepo(KafkaProducer *modules.KafkaProducer) *MailRepo {
	return &MailRepo{
		kp: KafkaProducer,
	}
}

func (r MailRepo) SendMail(mail *model.Mail) {
	mailData, _ := json.Marshal(&mail.Data)
	err := r.kp.Produce([]byte(mail.Key), mailData, mail.Topic)
	if err != nil {
		logger.Error("error while produce topic:" + err.Error())
	} else {
		logger.Info("successfully produce topic",
			zap.String("topic", mail.Topic),
			zap.String("key", mail.Key),
			zap.String("value", string(mailData)))
	}
}
