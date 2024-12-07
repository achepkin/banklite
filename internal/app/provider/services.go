package provider

import (
	"github.com/achepkin/banklite/internal/domain/services"
	"github.com/achepkin/banklite/internal/handlers"
)

func (p *Provider) AccountService() *services.AccountService {
	//return services.NewAccountService(p.AccountStorage())
	if p.accountService == nil {
		p.accountService = services.NewAccountService(p.AccountStorage())
	}

	return p.accountService
}

func (p *Provider) TransactionService() handlers.TransactionService {
	if p.transactionService == nil {
		p.transactionService = services.NewTransactionService(
			p.TransactionStorage(),
			p.AccountService(),
			p.TxValidator(),
		)
	}

	return p.transactionService
}
