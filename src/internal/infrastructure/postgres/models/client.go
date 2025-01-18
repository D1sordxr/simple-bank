package models

import (
	"github.com/google/uuid"
	"time"
)

type Client struct {
	ID        uuid.UUID
	FullName  string
	Email     string
	Status    string
	CreatedAt time.Time
}

type Phone struct {
	ID          uuid.UUID
	ClientID    uuid.UUID
	PhoneNumber string
	Country     int
	Code        int
	Number      int
	CreatedAt   time.Time
}
