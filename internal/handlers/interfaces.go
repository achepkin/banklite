package handlers

import (
	"context"
	"github.com/achepkin/banklite/internal/domain/entity"
)

type AccountService interface {
	CreateAccount(ctx context.Context, account *entity.Account) error
	GetAccount(ctx context.Context, id string) (*entity.Account, error)
	ListAccounts(ctx context.Context) ([]*entity.Account, error)
}

type TransactionService interface {
	CreateTransaction(ctx context.Context, accountID string, txType entity.TxType, amount float64) (*entity.Transaction, error)
	GetTransactions(ctx context.Context, accountID string) ([]*entity.Transaction, error)
	Transfer(ctx context.Context, fromAccountID string, toAccountID string, amount float64) (*entity.Transfer, error)
}
