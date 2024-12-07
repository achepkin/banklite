package services

import (
	"context"
	"fmt"
	"github.com/achepkin/banklite/internal/domain"
	"github.com/achepkin/banklite/internal/domain/entity"
	"time"
)

type Task struct {
	ctx       context.Context
	accountID string
	txType    entity.TxType
	amount    float64
}

type Result struct {
	txs []*entity.Transaction
	err error
}

type TransactionValidator interface {
	Validate(tx *entity.Transaction, account *entity.Account) error
}

type TransactionService struct {
	transactionRepo domain.TransactionRepository
	accountRepo     domain.AccountRepository
	validator       TransactionValidator
	txQueue         chan []Task
	txResult        chan Result
}

func NewTransactionService(
	transactionRepo domain.TransactionRepository,
	accountRepo domain.AccountRepository,
	validator TransactionValidator,
) *TransactionService {
	s := &TransactionService{
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
		validator:       validator,
		txQueue:         make(chan []Task),
		txResult:        make(chan Result),
	}

	s.StartTransactionProcessor()
	return s
}

func (s *TransactionService) StartTransactionProcessor() {

	process := func(tasks []Task) {
		txs := make([]*entity.Transaction, 0, len(tasks))
		for _, task := range tasks {
			tx, err := s.createTx(task.ctx, task.accountID, task.txType, task.amount)
			if err != nil {
				s.txResult <- Result{err: err}
				return
			}
			txs = append(txs, tx)
		}
		s.txResult <- Result{txs: txs}
	}

	go func() {
		for tt := range s.txQueue {
			process(tt)
		}
		close(s.txResult)
	}()
}

func (s *TransactionService) createTx(ctx context.Context, accountID string, txType entity.TxType, amount float64) (*entity.Transaction, error) {
	account, err := s.accountRepo.GetAccount(ctx, accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	tx := entity.NewTransaction(accountID, txType, amount)

	err = s.validator.Validate(tx, account)
	if err != nil {
		return nil, fmt.Errorf("failed to validate transaction: %w", err)
	}

	account.Apply(tx)

	err = s.transactionRepo.CreateTransaction(nil, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	err = s.accountRepo.UpdateAccount(account)
	if err != nil {
		return nil, fmt.Errorf("failed to update account: %w", err)
	}

	return tx, nil
}

func (s *TransactionService) CreateTransaction(ctx context.Context, accountID string, txType entity.TxType, amount float64) (*entity.Transaction, error) {
	s.txQueue <- []Task{
		{
			ctx:       ctx,
			accountID: accountID,
			txType:    txType,
			amount:    amount,
		},
	}
	result := <-s.txResult
	if result.err != nil {
		return nil, result.err
	}

	if len(result.txs) == 0 {
		return nil, fmt.Errorf("failed to create transaction")
	}
	return result.txs[0], nil
}

func (s *TransactionService) GetTransactions(ctx context.Context, accountID string) ([]*entity.Transaction, error) {
	_, err := s.accountRepo.GetAccount(ctx, accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}
	return s.transactionRepo.GetTransactions(ctx, accountID)
}

func (s *TransactionService) Transfer(ctx context.Context, fromAccountID string, toAccountID string, amount float64) (*entity.Transfer, error) {

	fmt.Printf("Transfering %f from %s to %s\n", amount, fromAccountID, toAccountID)
	s.txQueue <- []Task{
		{
			ctx:       ctx,
			accountID: fromAccountID,
			txType:    entity.TxTypeWithdrawal,
			amount:    amount,
		},
		{
			ctx:       ctx,
			accountID: toAccountID,
			txType:    entity.TxTypeDeposit,
			amount:    amount,
		},
	}
	result := <-s.txResult
	if result.err != nil {
		return nil, result.err
	}

	return &entity.Transfer{
		FromTransactionID: result.txs[0].ID,
		ToTransactionID:   result.txs[1].ID,
		Amount:            amount,
		Timestamp:         time.Now(),
	}, nil
}