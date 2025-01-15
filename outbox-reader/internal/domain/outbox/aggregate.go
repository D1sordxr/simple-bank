package outbox

import "time"

type Aggregate struct {
	OutboxID       string
	AggregateID    string
	AggregateType  string
	MessageType    string
	MessagePayload string
	Status         string
	CreatedAt      time.Time
}
