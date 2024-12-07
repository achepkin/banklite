package domain

import (
	"context"
	"github.com/achepkin/banklite/internal/domain/entity"
)

type AccountRepository interface {
	CreateAccount(context.Context, *entity.Account) error
	GetAccount(ctx context.Context, id string) (*entity.Account, error)
	ListAccounts(context.Context) ([]*entity.Account, error)
	UpdateAccount(account *entity.Account) error
}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) error
	GetTransactions(ctx context.Context, accountID string) ([]*entity.Transaction, error)
}
