package provider

import (
	"github.com/achepkin/banklite/internal/app"
	"github.com/achepkin/banklite/internal/domain/services"
	"github.com/achepkin/banklite/internal/domain/validators"
	"github.com/achepkin/banklite/internal/handlers"
	"github.com/achepkin/banklite/internal/storage"
)

type Provider struct {
	config *app.Config
	// handlers
	accountHandler     *handlers.AccountHandler
	transactionHandler *handlers.TransactionHandler
	// services
	accountService     *services.AccountService
	transactionService handlers.TransactionService
	// storages
	accountStorage     *storage.AccountStorageInMemory
	transactionStorage *storage.TransactionStorageInMemory
}

func (p *Provider) TransactionStorage() *storage.TransactionStorageInMemory {
	if p.transactionStorage == nil {
		p.transactionStorage = storage.NewTransactionStorageInMemory()
	}

	return p.transactionStorage
}

func (p *Provider) TxValidator() *validators.TransactionValidatorImpl {
	return validators.NewTransactionValidator()
}

func NewProvider(cfg *app.Config) *Provider {
	return &Provider{config: cfg}
}
