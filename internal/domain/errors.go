package domain

import "errors"

var (
	ErrAccountNotFound           = errors.New("account not found")
	ErrAccountExists             = errors.New("account already exists")
	ErrInvalidTransactionAmount  = errors.New("invalid transaction amount")
	ErrNotEnoughBalance          = errors.New("not enough balance")
	ErrTransactionAmountExceeded = errors.New("transaction amount exceeded")
)
