package service

import (
	"context"
	"fmt"

	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/pkg/entity"
	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/pkg/kafka"
)

type antiFraudService struct{}

func NewAntiFraudService() *antiFraudService {
	return &antiFraudService{}
}

func (af *antiFraudService) TransactionValidator(ctx context.Context, transaction entity.Transaction, producer *kafka.Producer) {
	transaction.TransactionStatus = entity.Approved

	if transaction.Value > 1000 {
		transaction.TransactionStatus = entity.Rejected
	}

	fmt.Println("ESTOU NO VALIDANDO E ESSE DEU ", transaction.TransactionStatus)

	err := producer.SendMessage(ctx, transaction.TransactionExternalID, transaction)
	if err != nil {
		fmt.Println(err)
		return
	}
}
