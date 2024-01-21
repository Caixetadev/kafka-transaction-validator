package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/postgresql"
	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/repository"
	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/routes"
	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/service"
	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/pkg/entity"
	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/pkg/kafka"
)

func main() {
	db, err := postgresql.New("postgres://root:password@localhost:5432/database")
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

	go func() {
		for {
			var transaction entity.Transaction

			consumer.ReadMessage(context.TODO(), &transaction)

			if len(transaction.TransactionStatus) > 0 {
				err := services.Update(context.TODO(), &transaction)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
	}()

	routes.NewTransactionRouter(services)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
