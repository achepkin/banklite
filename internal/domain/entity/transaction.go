package entity

import (
	"time"

	"github.com/google/uuid"
)

type TxType string

const (
	TxTypeDeposit    TxType = "deposit"
	TxTypeWithdrawal TxType = "withdrawal"
)

type Transaction struct {
	ID        string
	AccountID string
	Type      TxType // deposit or withdrawal
	Amount    float64
	Timestamp time.Time
}

func NewTransaction(accountID string, txType TxType, amount float64) *Transaction {
	return &Transaction{
		ID:        uuid.New().String(),
		AccountID: accountID,
		Type:      txType,
		Amount:    amount,
		Timestamp: time.Now(),
	}
}
