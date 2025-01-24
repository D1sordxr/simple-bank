package transaction

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres"
	converter "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/converters/transaction"
	"github.com/jackc/pgx/v5"
)

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

	model := converter.ConvertAggregateToModel(transaction)

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
