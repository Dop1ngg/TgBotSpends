package domain

import "time"

type TransactionType string

const (
	TxExpense TransactionType = "expense"
	TxIncome  TransactionType = "income"
)

type Transaction struct {
	ID          int64
	UserID      int64
	Type        TransactionType
	AmountCents int64
	Comment     string
	CreatedAt   time.Time
}
