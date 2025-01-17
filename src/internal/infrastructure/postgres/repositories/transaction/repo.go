package transaction

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres"
	eventConverter "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/converters/shared/event"
	outboxConverter "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/converters/shared/outbox"
	mainConverter "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/converters/transaction"
	"github.com/jackc/pgx/v5"
)

// TODO: implement methods

type Repository struct {
	Conn *postgres.Connection
}

func NewTransactionRepository(conn *postgres.Connection) *Repository {
	return &Repository{Conn: conn}
}

func (r *Repository) Create(ctx context.Context, tx interface{}, transaction transaction.Aggregate) error {
	conn := tx.(pgx.Tx)
	query := `INSERT INTO transactions (
                          id, 
                          source_account_id, 
                          destination_account_id,
                          currency,
                          amount,
                          status,
                          type,
                          description,
                          created_at
                          ) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`

	model := mainConverter.ConvertAggregateToModel(transaction)

	_, err := conn.Exec(ctx, query,
		model.ID,
		model.SourceAccountID,
		model.DestinationAccountID,
		model.Currency,
		model.Amount,
		model.Status,
		model.Type,
		model.Description,
		model.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
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

func (r *Repository) SaveOutboxEvent(ctx context.Context, tx interface{}, outbox outbox.Outbox) error {
	conn := tx.(pgx.Tx)
	query := `INSERT INTO outbox (
                          id, 
                          aggregate_id, 
                          aggregate_type,
                          message_type,
                          message_payload,
                          status,
                          created_at
                          ) 
				VALUES ($1, $2, $3, $4, $5, $6, $7);`

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
		return err
	}

	return nil
}
