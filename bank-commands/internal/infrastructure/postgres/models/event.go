package models

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID            uuid.UUID
	AggregateID   uuid.UUID
	AggregateType string
	EventType     string
	Payload       string
	CreatedAt     time.Time
}
