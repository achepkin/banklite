package validators

import (
	"github.com/achepkin/banklite/internal/domain"
	"github.com/achepkin/banklite/internal/domain/entity"
)

type TransactionValidatorImpl struct {
	maxAmount float64
}

func NewTransactionValidator() *TransactionValidatorImpl {
	return &TransactionValidatorImpl{
		maxAmount: 1000,
	}
}

func (v *TransactionValidatorImpl) Validate(transaction *entity.Transaction, account *entity.Account) error {
	if transaction.Amount <= 0 {
		return domain.ErrInvalidTransactionAmount
	}

	if transaction.Type == entity.TxTypeWithdrawal && account.Balance < transaction.Amount {
		return domain.ErrNotEnoughBalance
	}

	if transaction.Amount > v.maxAmount {
		return domain.ErrTransactionAmountExceeded
	}

	return nil
}
