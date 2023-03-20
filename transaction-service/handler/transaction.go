package handler

import (
	"encoding/json"
	"net/http"

	"github.com/danisbagus/go-messaging-kafka/transaction-service/dto"
	"github.com/danisbagus/go-messaging-kafka/transaction-service/service"
)

type TransactionHandler struct {
	service service.ITransactionService
}

func NewTransactionHanldler(service service.ITransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

func (rc TransactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	var request dto.NewTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := rc.service.Create(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	writeResponse(w, http.StatusCreated, data)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
