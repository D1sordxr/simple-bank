package queries

import "time"

type OutboxDTO struct {
	OutboxID       string
	AggregateID    string
	AggregateType  string
	MessageType    string
	MessagePayload string
	Status         string
	CreatedAt      time.Time
}

type OutboxDTOs []OutboxDTO
