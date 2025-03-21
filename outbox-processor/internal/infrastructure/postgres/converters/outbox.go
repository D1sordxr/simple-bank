package converters

import (
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/queries"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres/models"
)

func ConvertModelToDTO(model models.Outbox) queries.OutboxDTO {
	return queries.OutboxDTO{
		OutboxID:       model.ID.String(),
		AggregateID:    model.AggregateID.String(),
		AggregateType:  model.AggregateType,
		MessageType:    model.MessageType,
		MessagePayload: model.MessagePayload,
		Status:         model.Status,
		CreatedAt:      model.CreatedAt,
	}
}
