package event

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/account"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event/vo"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_vo"
	"time"
)

func NewAccountCreatedEvent(account account.Aggregate) (Event, error) {

	// TODO: ...

	eventID := sharedVO.NewUUID()
	// aggregateID := account.AccountID.String()
	aggregateType := vo.NewAccountAggregateType()
	eventType, err := vo.NewEventType(vo.TypeCreated)
	if err != nil {
		return Event{}, err
	}
	payload, err := vo.NewEventPayload(account)
	if err != nil {
		return Event{}, err
	}
	creationTime := time.Now()
	return Event{
		EventID: eventID,
		// AggregateID:   aggregateID,
		AggregateType: aggregateType,
		EventType:     eventType,
		Payload:       payload,
		CreatedAt:     creationTime,
	}, nil
}
