package models

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID        uuid.UUID
	ClientID  uuid.UUID
	Balance   float64
	Currency  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
