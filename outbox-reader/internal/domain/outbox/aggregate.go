package outbox

import "time"

type Aggregate struct {
	ID             string
	AggregateID    string
	AggregateType  string
	MessageType    string
	MessagePayload string
	Status         string
	CreatedAt      time.Time
	ProcessedAt    *time.Time
}
