package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/go-messaging-kafka/mail-service/model"
	"github.com/danisbagus/go-messaging-kafka/mail-service/modules"
	"github.com/danisbagus/go-messaging-kafka/mail-service/service"
	"github.com/segmentio/kafka-go"
)

type TransactionHandler struct {
	service service.ITransactionService
	kc      *modules.KafkaConsumer
}

func NewTransactionHandler(service service.ITransactionService, kc *modules.KafkaConsumer) *TransactionHandler {
	return &TransactionHandler{
		service: service,
		kc:      kc,
	}
}

func (h TransactionHandler) Consume(ctx context.Context, message kafka.Message) error {
	payload := model.TransactionPayload{}
	resMsg := ""

	err := json.Unmarshal(message.Value, &payload)
	if err != nil {
		return fmt.Errorf("failed unmarshal message value:" + err.Error())
	}

	switch payload.Action {
	case "create":
		err = h.service.SendTransactionPaid(payload.Data)
		if err != nil {
			return fmt.Errorf("failed send mail transaction paid:" + err.Error())
		}
		resMsg = "successfully send mail transaction paid"
	default:
		return fmt.Errorf("unsupported action:" + payload.Action)
	}

	if resMsg != "" {
		logger.Info(resMsg)
	}

	return nil
}
