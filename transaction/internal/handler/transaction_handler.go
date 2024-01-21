package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/service"
)

type TransactionService interface {
	Insert(ctx context.Context, transaction service.CreateTransactionInput) error
}

type TransactionHandler struct {
	TransactionService TransactionService
}

func (th *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/transaction" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var transaction service.CreateTransactionInput
	if err := json.Unmarshal(body, &transaction); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	go th.TransactionService.Insert(context.TODO(), transaction)

	w.Write([]byte("CRIADO COM SUCESSO"))
}
