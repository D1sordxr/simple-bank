package models

import (
	"github.com/google/uuid"
	"time"
)

type Client struct {
	ID         uuid.UUID `db:"id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	MiddleName string    `db:"middle_name"`
	Email      string    `db:"email"`
	Status     string    `db:"status"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type Phone struct {
	ID          uuid.UUID `db:"id"`
	ClientID    uuid.UUID `db:"client_id"`
	PhoneNumber string    `db:"phone_number"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
