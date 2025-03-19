package services

import (
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/commands"
	eventPkg "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/account"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/vo"
	sharedVO "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/shared_vo"
	"time"
)

type UpdateDomainSvc struct{}

func (UpdateDomainSvc) CreateUpdateEvent(upd account.UpdateEvent) (eventPkg.Event, error) {
	const op = "domain.account.UpdateDomainSvc.CreateUpdateEvent"

	aggID, err := sharedVO.NewUUIDFromString(upd.AccountID)
	if err != nil {
		return eventPkg.Event{}, fmt.Errorf("%s: %w", op, err)
	}

	eventType, err := vo.NewEventType(vo.TypeUpdated)
	if err != nil {
		return eventPkg.Event{}, fmt.Errorf("%s: %w", op, err)
	}

	payload, err := vo.NewEventPayload(upd)
	if err != nil {
		return eventPkg.Event{}, fmt.Errorf("%s: %w", op, err)
	}

	event := eventPkg.Event{
		EventID:       sharedVO.NewUUID(),
		AggregateID:   aggID,
		AggregateType: vo.NewAccountAggregateType(),
		EventType:     eventType,
		Payload:       payload,
		CreatedAt:     time.Now(),
	}

	return event, nil
}

func (UpdateDomainSvc) ConvertCommandToUpdEvent(c commands.UpdateAccountCommand) account.UpdateEvent {
	return account.UpdateEvent{
		AccountID:         c.AccountID,
		Amount:            c.Amount,
		BalanceUpdateType: c.BalanceUpdateType,
		Status:            c.Status,
		TransactionID:     c.TransactionID,
	}
}
