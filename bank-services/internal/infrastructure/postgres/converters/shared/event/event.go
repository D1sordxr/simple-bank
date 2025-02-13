package event

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/models"
)

func ConvertAggregateToModel(event event.Event) models.Event {
	return models.Event{
		ID:            event.EventID.Value,
		AggregateID:   event.AggregateID.Value,
		AggregateType: event.AggregateType.Type,
		EventType:     event.EventType.Type,
		Payload:       event.Payload.Payload,
		CreatedAt:     event.CreatedAt,
	}
}
