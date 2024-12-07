// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/achepkin/banklite/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

type TransactionRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *TransactionRepository) EXPECT() *TransactionRepository_Expecter {
	return &TransactionRepository_Expecter{mock: &_m.Mock}
}

// CreateTransaction provides a mock function with given fields: ctx, transaction
func (_m *TransactionRepository) CreateTransaction(ctx context.Context, transaction *entity.Transaction) error {
	ret := _m.Called(ctx, transaction)

	if len(ret) == 0 {
		panic("no return value specified for CreateTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Transaction) error); ok {
		r0 = rf(ctx, transaction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransactionRepository_CreateTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTransaction'
type TransactionRepository_CreateTransaction_Call struct {
	*mock.Call
}

// CreateTransaction is a helper method to define mock.On call
//   - ctx context.Context
//   - transaction *entity.Transaction
func (_e *TransactionRepository_Expecter) CreateTransaction(ctx interface{}, transaction interface{}) *TransactionRepository_CreateTransaction_Call {
	return &TransactionRepository_CreateTransaction_Call{Call: _e.mock.On("CreateTransaction", ctx, transaction)}
}

func (_c *TransactionRepository_CreateTransaction_Call) Run(run func(ctx context.Context, transaction *entity.Transaction)) *TransactionRepository_CreateTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Transaction))
	})
	return _c
}

func (_c *TransactionRepository_CreateTransaction_Call) Return(_a0 error) *TransactionRepository_CreateTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TransactionRepository_CreateTransaction_Call) RunAndReturn(run func(context.Context, *entity.Transaction) error) *TransactionRepository_CreateTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransactions provides a mock function with given fields: ctx, accountID
func (_m *TransactionRepository) GetTransactions(ctx context.Context, accountID string) ([]*entity.Transaction, error) {
	ret := _m.Called(ctx, accountID)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactions")
	}

	var r0 []*entity.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*entity.Transaction, error)); ok {
		return rf(ctx, accountID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*entity.Transaction); ok {
		r0 = rf(ctx, accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionRepository_GetTransactions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransactions'
type TransactionRepository_GetTransactions_Call struct {
	*mock.Call
}

// GetTransactions is a helper method to define mock.On call
//   - ctx context.Context
//   - accountID string
func (_e *TransactionRepository_Expecter) GetTransactions(ctx interface{}, accountID interface{}) *TransactionRepository_GetTransactions_Call {
	return &TransactionRepository_GetTransactions_Call{Call: _e.mock.On("GetTransactions", ctx, accountID)}
}

func (_c *TransactionRepository_GetTransactions_Call) Run(run func(ctx context.Context, accountID string)) *TransactionRepository_GetTransactions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *TransactionRepository_GetTransactions_Call) Return(_a0 []*entity.Transaction, _a1 error) *TransactionRepository_GetTransactions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TransactionRepository_GetTransactions_Call) RunAndReturn(run func(context.Context, string) ([]*entity.Transaction, error)) *TransactionRepository_GetTransactions_Call {
	_c.Call.Return(run)
	return _c
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
