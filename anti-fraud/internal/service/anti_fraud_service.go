package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/pkg/kafka"
)

type antiFraudService struct{}

func NewAntiFraudService() *antiFraudService {
	return &antiFraudService{}
}

type Transaction struct {
	TransactionExternalID string            `json:"transactionExternalId"`
	TransactionType       TransactionType   `json:"transactionType"`
	TransactionStatus     TransactionStatus `json:"transactionStatus"`
	Value                 float64           `json:"value"`
	CreatedAt             time.Time         `json:"createdAt"`
}

type TransactionType struct {
	Name string `json:"name"`
}

type TransactionStatus string

const (
	Approved TransactionStatus = "approved"
	Rejected TransactionStatus = "rejected"
)

func (af *antiFraudService) TransactionValidator(ctx context.Context, transaction Transaction, producer *kafka.Producer) {
	transaction.TransactionStatus = Approved

	if transaction.Value > 1000 {
		transaction.TransactionStatus = Rejected
	}

	err := producer.SendMessage(ctx, transaction.TransactionExternalID, transaction)
	if err != nil {
		fmt.Println(err)
		return
	}
}
