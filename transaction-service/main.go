package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/go-messaging-kafka/transaction-service/handler"
	"github.com/danisbagus/go-messaging-kafka/transaction-service/modules"
	"github.com/danisbagus/go-messaging-kafka/transaction-service/repo"
	"github.com/danisbagus/go-messaging-kafka/transaction-service/service"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		panic("error loading env file" + err.Error())
	}

	// create kafka topic
	newTopics := []string{
		os.Getenv("TOPIC_TRANSACTION"),
	}

	modules.CreateTopics(newTopics)

	// init kafka producers
	kp := modules.NewKafkaProducer()

	// multiplexer
	router := mux.NewRouter()

	// injenction
	mailRepo := repo.NewMailRepo(kp)
	transactionService := service.NewTransactionService(mailRepo)
	transactionHandler := handler.NewTransactionHanldler(transactionService)

	// routing
	router.HandleFunc("/api/transactions", transactionHandler.NewTransaction).Methods(http.MethodPost)

	// starting server
	port := "9050"
	logger.Info("Starting transaction service on port:" + port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", port), router))
}
