package transaction

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/converters"
	"github.com/jackc/pgx/v5"
)

// TODO: implement methods

type Repository struct {
	Conn *pgx.Conn
}

func NewTransactionRepository(conn *pgx.Conn) *Repository {
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

	model := converters.ConvertAggregateToModel(transaction)

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

func (r *Repository) SaveEvent(ctx context.Context, tx interface{}, transaction transaction.Aggregate) error {
	return nil
}

func (r *Repository) SaveOutboxEvent(ctx context.Context, tx interface{}, outbox outbox.Outbox) error {
	return nil
}
