package models

import (
	"github.com/google/uuid"
	"time"
)

type TransactionModel struct {
	ID                   uuid.UUID  `db:"id"`
	SourceAccountID      *uuid.UUID `db:"source_account_id"`
	DestinationAccountID *uuid.UUID `db:"destination_account_id"`
	Currency             string     `db:"currency"`
	Amount               float64    `db:"amount"`
	Status               string     `db:"status"`
	Type                 string     `db:"type"`
	Description          *string    `db:"description"`
	FailureReason        *string    `db:"failure_reason"`
	CreatedAt            time.Time  `db:"created_at"`
	UpdatedAt            time.Time  `db:"updated_at"`
}
