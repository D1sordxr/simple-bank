package event

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres"
	contextUtils "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/context-utils"
	eventConverter "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/converters/shared/event"
)

type Repository struct {
	Conn *postgres.Pool
}

func NewEventRepositoryV2(conn *postgres.Pool) *Repository {
	return &Repository{Conn: conn}
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

	tx, err := contextUtils.GetTxFromContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = tx.Exec(ctx, query,
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
