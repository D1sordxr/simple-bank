package outbox

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/outbox"
	outboxConverter "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/converters/shared/outbox"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/executor"
)

type Repository struct {
	Executor *executor.Executor
}

func NewOutboxRepository(executor *executor.Executor) *Repository {
	return &Repository{Executor: executor}
}

func (r *Repository) SaveOutboxEvent(ctx context.Context, outbox outbox.Outbox) error {
	const op = "postgres.OutboxRepository.SaveOutboxEvent"

	conn := r.Executor.GetExecutor(ctx)

	query := `INSERT INTO outbox (
        id, 
        aggregate_id, 
        aggregate_type,
        message_type,
        message_payload,
        status,
        created_at
    ) VALUES ($1, $2, $3, $4, $5, $6, $7);`

	model := outboxConverter.ConvertAggregateToModel(outbox)

	_, err := conn.Exec(ctx, query,
		model.ID,
		model.AggregateID,
		model.AggregateType,
		model.MessageType,
		model.MessagePayload,
		model.Status,
		model.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
