package models

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID             uuid.UUID `db:"id"`
	ClientID       uuid.UUID `db:"client_id"`
	AvailableMoney float64   `db:"available_money"`
	FrozenMoney    float64   `db:"frozen_money"`
	Currency       string    `db:"currency"`
	Status         string    `db:"status"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
