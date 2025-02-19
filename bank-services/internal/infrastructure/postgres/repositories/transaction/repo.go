package transaction

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction"
	converter "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/converters/transaction"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/executor"
)

type Repository struct {
	Executor *executor.Executor
}

func NewTransactionRepository(executor *executor.Executor) *Repository {
	return &Repository{Executor: executor}
}

func (r *Repository) Create(ctx context.Context, transaction transaction.Aggregate) error {
	const op = "postgres.TransactionRepository.Create"

	conn := r.Executor.GetExecutor(ctx)

	query := `INSERT INTO transactions (
        id, 
        source_account_id, 
        destination_account_id,
        currency,
        amount,
        status,
        type,
        description,
        failure_reason,
        created_at,
        updated_at
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`

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
		model.FailureReason,
		model.CreatedAt,
		model.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrFailedToCreateTransaction)
	}

	return nil
}
