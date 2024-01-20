package entity

import "time"

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
	Pending  TransactionStatus = "pending"
	Approved TransactionStatus = "approved"
	Rejected TransactionStatus = "rejected"
)
