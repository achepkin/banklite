package services

import (
	"context"
	"github.com/achepkin/banklite/internal/domain"
	"github.com/achepkin/banklite/internal/domain/entity"
)

type AccountService struct {
	accRepo domain.AccountRepository
}

func NewAccountService(accRepo domain.AccountRepository) *AccountService {
	return &AccountService{accRepo: accRepo}
}

func (s AccountService) CreateAccount(ctx context.Context, account *entity.Account) error {
	return s.accRepo.CreateAccount(nil, account)
}

func (s AccountService) GetAccount(ctx context.Context, id string) (*entity.Account, error) {
	return s.accRepo.GetAccount(nil, id)
}

func (s AccountService) ListAccounts(context.Context) ([]*entity.Account, error) {
	return s.accRepo.ListAccounts(nil)
}

func (s AccountService) UpdateAccount(account *entity.Account) error {
	return s.accRepo.UpdateAccount(account)
}
