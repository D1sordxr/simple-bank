package outbox

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/models"
)

func ConvertAggregateToModel(outbox outbox.Outbox) models.Outbox {
	return models.Outbox{
		ID:             outbox.OutboxID.Value,
		AggregateID:    outbox.AggregateID.Value,
		AggregateType:  outbox.AggregateType.Type,
		MessageType:    outbox.MessageType.Type,
		MessagePayload: outbox.MessagePayload.Payload,
		Status:         outbox.Status.Status,
		CreatedAt:      outbox.CreatedAt,
	}
}
