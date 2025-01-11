package models

import (
	"github.com/google/uuid"
	"time"
)

type TransactionModel struct {
	ID                   uuid.UUID
	SourceAccountID      *uuid.UUID
	DestinationAccountID *uuid.UUID
	Currency             string
	Amount               float64
	Status               string
	Type                 string
	Description          *string
	CreatedAt            time.Time
}
