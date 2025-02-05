package models

import (
	"github.com/google/uuid"
	"time"
)

type Outbox struct {
	ID             uuid.UUID
	AggregateID    uuid.UUID
	AggregateType  string
	MessageType    string
	MessagePayload string
	Status         string
	CreatedAt      time.Time
}
