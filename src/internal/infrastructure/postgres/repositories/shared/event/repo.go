package event

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres"
	eventConverter "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/converters/shared/event"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn *postgres.Connection
}

func NewEventRepository(conn *postgres.Connection) *Repository {
	return &Repository{Conn: conn}
}

func (r *Repository) SaveEvent(ctx context.Context, tx interface{}, event event.Event) error {
	conn := tx.(pgx.Tx)
	query := `INSERT INTO events (
                          id, 
                          aggregate_id, 
                          aggregate_type,
                          event_type,
                          payload,
                          created_at
                          ) 
				VALUES ($1, $2, $3, $4, $5, $6);`

	model := eventConverter.ConvertAggregateToModel(event)

	_, err := conn.Exec(ctx, query,
		model.ID,
		model.AggregateID,
		model.AggregateType,
		model.EventType,
		model.Payload,
		model.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
