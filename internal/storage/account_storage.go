package storage

import (
	"context"
	"sync"

	"github.com/achepkin/banklite/internal/domain"
	"github.com/achepkin/banklite/internal/domain/entity"
)

var _ domain.AccountRepository = (*AccountStorageInMemory)(nil)

type AccountStorageInMemory struct {
	accounts map[string]*entity.Account
	mu       sync.RWMutex
}

func NewAccountStorageInMemory() *AccountStorageInMemory {
	return &AccountStorageInMemory{
		accounts: make(map[string]*entity.Account),
	}
}

func (a *AccountStorageInMemory) CreateAccount(ctx context.Context, account *entity.Account) error {
	if _, ok := a.accounts[account.ID]; ok {
		return domain.ErrAccountExists
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.accounts[account.ID] = account
	return nil
}

func (a *AccountStorageInMemory) GetAccount(ctx context.Context, id string) (*entity.Account, error) {
	if account, ok := a.accounts[id]; ok {
		return account, nil
	}
	return nil, domain.ErrAccountNotFound
}

func (a *AccountStorageInMemory) ListAccounts(context.Context) ([]*entity.Account, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	accounts := make([]*entity.Account, 0, len(a.accounts))
	for _, account := range a.accounts {
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (a *AccountStorageInMemory) UpdateAccount(account *entity.Account) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.accounts[account.ID]; !ok {
		return domain.ErrAccountNotFound
	}
	a.accounts[account.ID] = account
	return nil
}
