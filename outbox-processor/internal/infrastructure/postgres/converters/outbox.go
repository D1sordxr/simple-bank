package converters

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/outbox"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/models"
)

func ConvertModelToAggregate(model models.Outbox) outbox.Aggregate {
	return outbox.Aggregate{
		OutboxID:       model.ID.String(),
		AggregateID:    model.AggregateID.String(),
		AggregateType:  model.AggregateType,
		MessageType:    model.MessageType,
		MessagePayload: model.MessagePayload,
		Status:         model.Status,
		CreatedAt:      model.CreatedAt,
	}
}
