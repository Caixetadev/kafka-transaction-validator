package main

import (
	"context"

	"github.com/Caixetadev/fraud-check-kafka-integration/anti-fraud/internal/service"

	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/pkg/kafka"
)

func main() {
	services := service.NewAntiFraudService()

	producer := kafka.NewProducer([]string{"localhost:9092"}, "VALIDATED_TRANSACTION")
	defer producer.Close()

	consumer := kafka.NewConsumer([]string{"localhost:9092"}, "CREATED_TRANSACTION")
	defer consumer.Close()

	go func() {
		for {
			var transaction service.Transaction

			consumer.ReadMessage(context.TODO(), &transaction)

			if transaction.TransactionStatus == "pending" {
				services.TransactionValidator(context.TODO(), transaction, producer)
			}
		}
	}()

	select {}
}
