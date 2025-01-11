package outbox

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/models"
)

func ConvertAggregateToModel(outbox outbox.Outbox) models.Outbox {
	return models.Outbox{
		OutboxID:       outbox.OutboxID.Value,
		AggregateID:    outbox.AggregateID.Value,
		AggregateType:  outbox.AggregateType.Type,
		MessageType:    outbox.MessageType.Type,
		MessagePayload: outbox.MessagePayload.String(),
		Status:         outbox.Status.Status,
		CreatedAt:      outbox.CreatedAt,
	}
}
