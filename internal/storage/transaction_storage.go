package storage

import (
	"context"
	"sync"

	"github.com/achepkin/banklite/internal/domain/entity"
)

type TransactionStorageInMemory struct {
	transactions map[string]*entity.Transaction
	mu           sync.RWMutex
}

func NewTransactionStorageInMemory() *TransactionStorageInMemory {
	return &TransactionStorageInMemory{
		transactions: make(map[string]*entity.Transaction),
	}
}

func (t *TransactionStorageInMemory) CreateTransaction(ctx context.Context, transaction *entity.Transaction) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.transactions[transaction.ID] = transaction
	return nil
}

func (t *TransactionStorageInMemory) GetTransactions(ctx context.Context, accountID string) ([]*entity.Transaction, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	transactions := make([]*entity.Transaction, 0)
	for _, transaction := range t.transactions {
		if transaction.AccountID == accountID {
			transactions = append(transactions, transaction)
		}
	}
	return transactions, nil
}
