package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Caixetadev/fraud-check-kafka-integration/anti-fraud/internal/service"

	kafkac "github.com/Caixetadev/fraud-check-kafka-integration/transaction/pkg/kafka"
	"github.com/segmentio/kafka-go"
)

func main() {
	services := service.NewAntiFraudService()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "CREATED_TRANSACTION",
		Partition: 0,
		MaxBytes:  10e6,
	})

	producer := kafkac.NewProducer([]string{"localhost:9092"}, "VALIDATED_TRANSACTION")
	defer producer.Close()

	go func() {
		for {
			var transaction service.Transaction

			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Println("Error reading message:", err)
				continue
			}

			err = json.Unmarshal(m.Value, &transaction)
			if err != nil {
				log.Println("Error decoding JSON:", err)
				continue
			}

			services.TransactionValidator(context.TODO(), transaction, producer)
		}
	}()

	select {}
}
