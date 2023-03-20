package main

import (
	"context"
	"fmt"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/go-messaging-kafka/mail-service/handler"
	"github.com/danisbagus/go-messaging-kafka/mail-service/modules"
	"github.com/danisbagus/go-messaging-kafka/mail-service/repo"
	"github.com/danisbagus/go-messaging-kafka/mail-service/service"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		panic("error loading env file" + err.Error())
	}

	ctx := context.Background()

	smtp := modules.NewSmtp()
	kc := modules.NewKafkaConsumer()

	defer kc.Reader.Close()

	transactionRepo := repo.NewTransactionRepo(smtp)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService, kc)

	mapHandler := map[string]modules.ConsumeHandler{
		"topic_transaction": transactionHandler.Consume,
	}

	forever := make(chan bool)

	go func() {
		for {
			message, err := kc.Reader.ReadMessage(ctx)
			if err != nil {
				logger.Error(fmt.Sprintf("%s failed reade message: ", err))
				break
			}

			logger.Info("starting consume topic",
				zap.String("topic", message.Topic),
				zap.String("key", string(message.Key)),
				zap.String("value", string(message.Value)),
			)

			handler, ok := mapHandler[message.Topic]
			if !ok {
				logger.Error("undefined consumer handler. topic: " + message.Topic)
				break
			}

			switch message.Topic {
			case "topic_transaction":
				err = handler(ctx, message)
			default:
				err = fmt.Errorf("unsuported topic:" + message.Topic)
			}

			if err != nil {
				logger.Error(err.Error())
			}
		}
	}()

	logger.Info(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}
