package outbox

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres"
	outboxConverter "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/converters/shared/outbox"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn *postgres.Connection
}

func NewOutboxRepository(conn *postgres.Connection) *Repository {
	return &Repository{Conn: conn}
}

func (r *Repository) SaveOutboxEvent(ctx context.Context, tx interface{}, outbox outbox.Outbox) error {
	const op = "postgres.OutboxRepository.SaveOutboxEvent"

	conn := tx.(pgx.Tx)
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
