package routes

import (
	"net/http"

	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/handler"
	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/service"
)

func NewTransactionRouter(service *service.TransactionService) {
	handler := handler.TransactionHandler{
		TransactionService: service,
	}

	http.HandleFunc("/transaction", handler.CreateTransaction)
}
