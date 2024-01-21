package main

import (
	"context"
	"fmt"

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
	transactionInput := service.CreateTransactionInput{
		AccountExternalIDDebit:  "teste",
		Value:                   230.0,
		TransferTypeID:          1,
		AccountExternalIDCredit: "teste",
	}

	producer := kafka.NewProducer([]string{"localhost:9092"}, "CREATED_TRANSACTION")
	defer producer.Close()

	service := service.NewTransactionService(repository, producer)

	err = service.Insert(context.TODO(), transactionInput)
	if err != nil {
		fmt.Println(err)
	}
}
