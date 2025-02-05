package event

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event/vo"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction"
	"time"
)

func NewTransactionCreatedEvent(tx transaction.Aggregate) (Event, error) {
	eventID := sharedVO.NewUUID()
	aggregateID := tx.TransactionID
	aggregateType := vo.NewTransactionAggregateType()
	eventType, err := vo.NewEventType(vo.TypeCreated)
	if err != nil {
		return Event{}, err
	}
	payload, err := vo.NewEventPayload(tx)
	if err != nil {
		return Event{}, err
	}
	creationTime := time.Now()
	return Event{
		EventID:       eventID,
		AggregateID:   aggregateID,
		AggregateType: aggregateType,
		EventType:     eventType,
		Payload:       payload,
		CreatedAt:     creationTime,
	}, nil
}
