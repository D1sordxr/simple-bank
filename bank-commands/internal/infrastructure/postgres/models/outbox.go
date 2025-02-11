package models

import (
	"github.com/google/uuid"
	"time"
)

type Outbox struct {
	ID             uuid.UUID `db:"id"`
	AggregateID    uuid.UUID `db:"aggregate_id"`
	AggregateType  string    `db:"aggregate_type"`
	MessageType    string    `db:"message_type"`
	MessagePayload string    `db:"message_payload"`
	Status         string    `db:"status"`
	CreatedAt      time.Time `db:"created_at"`
}
