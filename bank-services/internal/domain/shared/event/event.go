package event

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/vo"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/outbox"
	vo2 "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/outbox/vo"
	sharedVO "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/shared_vo"
	"time"
)

type Event struct {
	EventID       sharedVO.UUID    // Event unique ID
	AggregateID   sharedVO.UUID    // References to aggregate unique ID
	AggregateType vo.AggregateType // Client, Account, Transaction
	EventType     vo.EventType     // Created, Updated, Deleted
	Payload       vo.EventPayload  // Contains marshalled JSON
	CreatedAt     time.Time        // Creation time
}

func (e *Event) OutboxFromEvent() outbox.Outbox {
	status := vo2.OutboxStatus{Status: vo2.StatusPending}
	return outbox.Outbox{
		OutboxID:       sharedVO.NewUUID(),
		AggregateID:    e.EventID,
		AggregateType:  e.AggregateType,
		MessageType:    e.EventType,
		MessagePayload: e.Payload,
		Status:         status,
		CreatedAt:      time.Now(),
	}
}
