package entity

import "time"

type Transfer struct {
	FromTransactionID string
	ToTransactionID   string
	Amount            float64
	Timestamp         time.Time
}
