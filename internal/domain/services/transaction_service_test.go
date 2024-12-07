package services

import (
	"context"
	mocks "github.com/achepkin/banklite/internal/pkg/mocks/domain"
	mocks_services "github.com/achepkin/banklite/internal/pkg/mocks/domain/services"
	"testing"

	"github.com/achepkin/banklite/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransactionService_CreateTransaction(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	mockAccountRepo := new(mocks.AccountRepository)
	mockValidator := new(mocks_services.TransactionValidator)

	service := NewTransactionService(mockTransactionRepo, mockAccountRepo, mockValidator)

	ctx := context.Background()
	accountID := "test-account-id"
	txType := entity.TxTypeDeposit
	amount := 100.0

	account := &entity.Account{ID: accountID, Balance: 1000.0}

	mockAccountRepo.EXPECT().GetAccount(ctx, accountID).Return(account, nil)
	mockValidator.EXPECT().Validate(mock.Anything, account).Return(nil)
	mockTransactionRepo.EXPECT().CreateTransaction(mock.Anything, mock.Anything).Return(nil)
	mockAccountRepo.EXPECT().UpdateAccount(mock.Anything).Return(nil)

	tx, err := service.CreateTransaction(ctx, accountID, txType, amount)

	assert.NoError(t, err)
	assert.NotNil(t, tx)
	assert.Equal(t, accountID, tx.AccountID)
	assert.Equal(t, txType, tx.Type)
	assert.Equal(t, amount, tx.Amount)

	mockAccountRepo.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactions(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	mockAccountRepo := new(mocks.AccountRepository)
	mockValidator := new(mocks_services.TransactionValidator)

	service := NewTransactionService(mockTransactionRepo, mockAccountRepo, mockValidator)

	ctx := context.Background()
	accountID := "test-account-id"

	account := &entity.Account{ID: accountID, Balance: 1000.0}
	tx1 := entity.NewTransaction(accountID, entity.TxTypeDeposit, 100.0)
	tx2 := entity.NewTransaction(accountID, entity.TxTypeWithdrawal, 50.0)
	transactions := []*entity.Transaction{
		tx1,
		tx2,
	}

	mockAccountRepo.EXPECT().GetAccount(ctx, accountID).Return(account, nil).Once()
	mockTransactionRepo.EXPECT().GetTransactions(ctx, accountID).Return(transactions, nil).Once()

	result, err := service.GetTransactions(ctx, accountID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, transactions, result)

	mockAccountRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}

func TestTransactionService_Transfer(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	mockAccountRepo := new(mocks.AccountRepository)
	mockValidator := new(mocks_services.TransactionValidator)

	service := NewTransactionService(mockTransactionRepo, mockAccountRepo, mockValidator)

	ctx := context.Background()
	fromAccountID := "from-account-id"
	toAccountID := "to-account-id"
	amount := 100.0

	fromAccount := &entity.Account{ID: fromAccountID, Balance: 1000.0}
	toAccount := &entity.Account{ID: toAccountID, Balance: 500.0}

	mockAccountRepo.EXPECT().GetAccount(ctx, fromAccountID).Return(fromAccount, nil).Once()
	mockAccountRepo.EXPECT().GetAccount(ctx, toAccountID).Return(toAccount, nil).Once()
	mockValidator.EXPECT().Validate(mock.Anything, fromAccount).Return(nil).Once()
	mockValidator.EXPECT().Validate(mock.Anything, toAccount).Return(nil).Once()
	mockTransactionRepo.EXPECT().CreateTransaction(nil, mock.Anything).Return(nil).Twice()
	mockAccountRepo.EXPECT().UpdateAccount(fromAccount).Return(nil).Once()
	mockAccountRepo.EXPECT().UpdateAccount(toAccount).Return(nil).Once()

	transfer, err := service.Transfer(ctx, fromAccountID, toAccountID, amount)

	assert.NoError(t, err)
	assert.NotNil(t, transfer)
	assert.Equal(t, amount, transfer.Amount)

	mockAccountRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}
