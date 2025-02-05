package outbox

import "time"

const (
	BatchSize = 3
)

type Aggregate struct {
	OutboxID       string
	AggregateID    string
	AggregateType  string
	MessageType    string
	MessagePayload string
	Status         string
	CreatedAt      time.Time
}
