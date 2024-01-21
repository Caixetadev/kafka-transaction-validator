package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/entity"
	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/postgresql"
	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/repository"
	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/service"
	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/pkg/kafka"
)

func main() {
	db, err := postgresql.New("postgres://root:password@localhost:5432/transaction")
	if err != nil {
		fmt.Println(err)
		return
	}

	repository := repository.NewTransactionRepository(db)

	producer := kafka.NewProducer([]string{"localhost:9092"}, "CREATED_TRANSACTION")
	defer producer.Close()

	consumer := kafka.NewConsumer([]string{"localhost:9092"}, "VALIDATED_TRANSACTION")
	defer consumer.Close()

	services := service.NewTransactionService(repository, producer)

	http.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/transaction" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}

		if r.Method != "POST" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
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

		go services.Insert(context.TODO(), transaction)

		w.Write([]byte("CRIADO COM SUCESSO"))
	})

	go func() {
		for {
			var transaction entity.Transaction

			consumer.ReadMessage(context.TODO(), &transaction)

			if len(transaction.TransactionStatus) > 0 {
				services.Update(context.TODO(), &transaction)
			}
		}
	}()

	http.ListenAndServe(":8080", nil)
}
