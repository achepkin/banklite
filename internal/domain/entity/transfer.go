package entity

import "time"

type Transfer struct {
	FromTransactionID string
	ToTransactionID   string
	Amount            float64
	Timestamp         time.Time
}

func NewTransfer(fromTransactionID, toTransactionID string, amount float64) (*Transfer, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}
	return &Transfer{
		FromTransactionID: fromTransactionID,
		ToTransactionID:   toTransactionID,
		Amount:            amount,
		Timestamp:         time.Now(),
	}, nil
}
