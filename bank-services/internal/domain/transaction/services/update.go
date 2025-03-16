package services

import (
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/commands"
	eventPkg "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/vo"
	sharedVO "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/shared_vo"
	"time"
)

type UpdateDomainSvc struct{}

func (UpdateDomainSvc) CreateUpdateEvent(c commands.UpdateTransactionCommand) (eventPkg.Event, error) {
	const op = "domain.account.UpdateDomainSvc.CreateUpdateEvent"

	aggID, err := sharedVO.NewUUIDFromString(c.TransactionID)
	if err != nil {
		return eventPkg.Event{}, fmt.Errorf("%s: %w", op, err)
	}

	eventType, err := vo.NewEventType(vo.TypeUpdated)
	if err != nil {
		return eventPkg.Event{}, fmt.Errorf("%s: %w", op, err)
	}

	payload, err := vo.NewEventPayload(c)
	if err != nil {
		return eventPkg.Event{}, fmt.Errorf("%s: %w", op, err)
	}

	event := eventPkg.Event{
		EventID:       sharedVO.NewUUID(),
		AggregateID:   aggID,
		AggregateType: vo.NewTransactionAggregateType(),
		EventType:     eventType,
		Payload:       payload,
		CreatedAt:     time.Now(),
	}

	return event, nil
}
