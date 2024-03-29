package service

import (
	"context"
	"time"

	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/pkg/entity"
	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/pkg/kafka"
	"github.com/google/uuid"
)

func NewTransaction(input CreateTransactionInput) *entity.Transaction {
	return &entity.Transaction{
		TransactionExternalID: uuid.New().String(),
		TransactionType: entity.TransactionType{
			Name: "ExampleType",
		},
		TransactionStatus: entity.Pending,
		Value:             input.Value,
		CreatedAt:         time.Now(),
	}
}

type CreateTransactionInput struct {
	AccountExternalIDDebit  string  `json:"accountExternalIdDebit"`
	AccountExternalIDCredit string  `json:"accountExternalIdCredit"`
	TransferTypeID          int     `json:"transferTypeId"`
	Value                   float64 `json:"value"`
}

type TransactionRepository interface {
	Insert(ctx context.Context, transaction *entity.Transaction) error
	Update(ctx context.Context, transaction *entity.Transaction) error
}

type TransactionService struct {
	transactionRepository TransactionRepository
	producer              *kafka.Producer
}

func NewTransactionService(transactionRepository TransactionRepository, producer *kafka.Producer) *TransactionService {
	return &TransactionService{
		transactionRepository: transactionRepository,
		producer:              producer,
	}
}

func (ts *TransactionService) Insert(ctx context.Context, transactionInput CreateTransactionInput) error {
	transaction := NewTransaction(transactionInput)

	err := ts.transactionRepository.Insert(ctx, transaction)
	if err != nil {
		return err
	}

	return ts.producer.SendMessage(ctx, transaction.TransactionExternalID, transaction)
}

func (ts *TransactionService) Update(ctx context.Context, transaction *entity.Transaction) error {
	return ts.transactionRepository.Update(ctx, transaction)
}
