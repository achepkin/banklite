package provider

import "github.com/achepkin/banklite/internal/storage"

func (p *Provider) AccountStorage() *storage.AccountStorageInMemory {
	if p.accountStorage == nil {
		p.accountStorage = storage.NewAccountStorageInMemory()
	}

	return p.accountStorage
}
