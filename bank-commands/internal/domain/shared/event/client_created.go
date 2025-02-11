package event

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/client"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event/vo"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_vo"
	"time"
)

func NewClientCreatedEvent(client client.Aggregate) (Event, error) {
	eventID := sharedVO.NewUUID()
	aggregateID := client.ClientID
	aggregateType := vo.NewClientAggregateType()
	eventType, err := vo.NewEventType(vo.TypeCreated)
	if err != nil {
		return Event{}, err
	}
	payload, err := vo.NewEventPayload(client)
	if err != nil {
		return Event{}, err
	}
	return Event{
		EventID:       eventID,
		AggregateID:   aggregateID,
		AggregateType: aggregateType,
		EventType:     eventType,
		Payload:       payload,
		CreatedAt:     time.Now(),
	}, nil
}
