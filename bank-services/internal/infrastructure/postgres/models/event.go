package models

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID            uuid.UUID `db:"id"`
	AggregateID   uuid.UUID `db:"aggregate_id"`
	AggregateType string    `db:"aggregate_type"`
	EventType     string    `db:"event_type"`
	Payload       string    `db:"payload"`
	CreatedAt     time.Time `db:"created_at"`
}
