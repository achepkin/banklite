package provider

import "github.com/achepkin/banklite/internal/handlers"

func (p *Provider) AccountHandler() *handlers.AccountHandler {
	//return handlers.NewAccountHandler(p.AccountService())
	if p.accountHandler == nil {
		p.accountHandler = handlers.NewAccountHandler(p.AccountService())
	}

	return p.accountHandler
}
func (p *Provider) TransactionHandler() *handlers.TransactionHandler {
	if p.transactionHandler == nil {
		p.transactionHandler = handlers.NewTransactionHandler(p.TransactionService())
	}

	return p.transactionHandler
}
