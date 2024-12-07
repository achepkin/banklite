package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type TxType string

var (
	ErrInvalidAmount = errors.New("invalid amount")
)

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
