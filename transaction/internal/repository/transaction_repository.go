package repository

import (
	"context"

	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/entity"
	"github.com/Caixetadev/fraud-check-kafka-integration/transaction/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ service.TransactionRepository = (*transactionRepository)(nil)

type transactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *transactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (tr *transactionRepository) Insert(ctx context.Context, transaction *entity.Transaction) error {
	query := "INSERT INTO transactions (id, type, status, value, created_at) VALUES ($1, $2, $3, $4, $5)"

	_, err := tr.db.Exec(
		ctx,
		query,
		transaction.TransactionExternalID,
		transaction.TransactionType.Name,
		transaction.TransactionStatus,
		transaction.Value,
		transaction.CreatedAt,
	)

	return err
}

func (tr *transactionRepository) Update(ctx context.Context, transaction *entity.Transaction) error {
	query := "UPDATE transactions SET status = $1 WHERE id = $2"

	_, err := tr.db.Exec(ctx, query, transaction.TransactionStatus, transaction.TransactionExternalID)

	return err
}
