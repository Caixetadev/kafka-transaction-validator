package service

import (
	"context"
	"time"

	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/entity"
)

func NewTransaction(input CreateTransactionInput) *entity.Transaction {
	return &entity.Transaction{
		TransactionExternalID: "teste123",
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
}

type transactionService struct {
	transactionRepository TransactionRepository
}

func NewTransactionService(transactionRepository TransactionRepository) *transactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
	}
}

func (ts *transactionService) Insert(ctx context.Context, transactionInput CreateTransactionInput) error {
	transaction := NewTransaction(transactionInput)

	err := ts.transactionRepository.Insert(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}
