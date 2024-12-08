package entity

import "time"

type Account struct {
	ID        string
	Owner     string
	Balance   float64
	CreatedAt time.Time
}

func (a *Account) Apply(tx *Transaction) {
	switch tx.Type {
	case TxTypeDeposit:
		a.Balance += tx.Amount
	case TxTypeWithdrawal:
		a.Balance -= tx.Amount
	}
}

func (a *Account) Rollback(tx *Transaction) {
	switch tx.Type {
	case TxTypeDeposit:
		a.Balance -= tx.Amount
	case TxTypeWithdrawal:
		a.Balance += tx.Amount
	}
}
