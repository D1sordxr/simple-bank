package event

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	eventConverter "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/converters/shared/event"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/executor"
)

type Repository struct {
	Executor *executor.Executor
}

func NewEventRepositoryV2(executor *executor.Executor) *Repository {
	return &Repository{Executor: executor}
}

func (r *Repository) SaveEvent(ctx context.Context, event event.Event) error {
	const op = "postgres.EventRepository.SaveEvent"
	query := `INSERT INTO events (
        id, 
        aggregate_id, 
        aggregate_type,
        event_type,
        payload,
        created_at
    ) VALUES ($1, $2, $3, $4, $5, $6);`

	model := eventConverter.ConvertAggregateToModel(event)

	conn := r.Executor.GetExecutor(ctx)

	_, err := conn.Exec(ctx, query,
		model.ID,
		model.AggregateID,
		model.AggregateType,
		model.EventType,
		model.Payload,
		model.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrFailedEventCreation)
	}

	return nil
}
