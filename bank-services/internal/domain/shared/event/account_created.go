package event

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/account"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/vo"
	sharedVO "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/shared_vo"
	"time"
)

func NewAccountCreatedEvent(account account.Aggregate) (Event, error) {
	eventID := sharedVO.NewUUID()
	aggregateID := account.AccountID
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
		EventID:       eventID,
		AggregateID:   aggregateID,
		AggregateType: aggregateType,
		EventType:     eventType,
		Payload:       payload,
		CreatedAt:     creationTime,
	}, nil
}
