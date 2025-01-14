package outbox

import "time"

// TODO: outbox

type Outbox struct {
	ID             string
	AggregateID    string
	AggregateType  string
	MessageType    string
	MessagePayload string
	Status         string
	CreatedAt      time.Time
	ProcessedAt    *time.Time
}

//type Outbox struct {
//	OutboxID       sharedVO.UUID         // Outbox unique ID
//	AggregateID    sharedVO.UUID         // References to aggregate unique ID
//	AggregateType  eventVO.AggregateType // Client, Account, Transaction
//	MessageType    eventVO.EventType     // Created, Updated, Deleted
//	MessagePayload eventVO.EventPayload  // Contains marshalled JSON
//	Status         vo.OutboxStatus       // Pending, Processed, Failed
//	CreatedAt      time.Time             // Creation time
//}
//
//func NewOutboxEvent(event event.Event) (Outbox, error) {
//	status, err := vo.NewOutboxStatus(vo.StatusPending)
//	if err != nil {
//		return Outbox{}, err
//	}
//	return Outbox{
//		OutboxID:       sharedVO.NewUUID(),
//		AggregateID:    event.AggregateID,
//		AggregateType:  event.AggregateType,
//		MessageType:    event.EventType,
//		MessagePayload: event.Payload,
//		Status:         status,
//		CreatedAt:      time.Now(),
//	}, nil
//}
