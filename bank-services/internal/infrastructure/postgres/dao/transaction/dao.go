package transaction

import (
	"context"
	"errors"
	"fmt"
	"github.com/D1sordxr/packages/postgres/executor"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/models"
	"github.com/jackc/pgx/v5"
)

type DAO struct {
	Executor *executor.Manager
}

func NewDAO(exec *executor.Manager) *DAO {
	return &DAO{Executor: exec}
}

func (d *DAO) GetTransaction(ctx context.Context, id string) (model models.TransactionModel, err error) {
	const op = "postgres.TransactionDAO.GetTransaction"

	conn := d.Executor.GetExecutor(ctx)

	query := `
		SELECT id, 
        source_account_id, 
        destination_account_id,
        currency,
        amount,
        status,
        type,
        description,
        failure_reason,
        created_at,
        updated_at FROM transactions WHERE id = $1
	`

	err = conn.QueryRow(ctx, query, id).Scan(
		&model.ID, &model.SourceAccountID, &model.DestinationAccountID, &model.Currency, &model.Amount,
		&model.Status, &model.Type, &model.Description, &model.FailureReason, &model.CreatedAt, &model.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.TransactionModel{}, fmt.Errorf("%s: %w: %w", op, ErrTransactionNotFound, err)
		}

		return models.TransactionModel{}, fmt.Errorf("%s: %w: %w", op, ErrReadingTransaction, err)
	}

	return model, nil
}

func (d *DAO) GetProjection(ctx context.Context, id string) (model models.TransactionModel, err error) {
	const op = "postgres.TransactionDAO.GetProjection"

	conn := d.Executor.GetExecutor(ctx)

	query := `
		SELECT id, 
        status,
        failure_reason FROM transactions WHERE id = $1
	`

	err = conn.QueryRow(ctx, query, id).Scan(
		&model.ID, &model.Status, &model.FailureReason,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.TransactionModel{}, fmt.Errorf("%s: %w: %w", op, ErrTransactionNotFound, err)
		}

		return models.TransactionModel{}, fmt.Errorf("%s: %w: %w", op, ErrReadingTransaction, err)
	}

	return model, nil
}

func (d *DAO) UpdateProjection(ctx context.Context, model models.TransactionModel) error {
	const op = "postgres.TransactionDAO.UpdateProjection"

	conn := d.Executor.GetExecutor(ctx)

	query := `
		UPDATE transactions SET status = $1, failure_reason = $2 WHERE id = $3
	`

	_, err := conn.Exec(ctx, query, model.Status, model.FailureReason, model.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
