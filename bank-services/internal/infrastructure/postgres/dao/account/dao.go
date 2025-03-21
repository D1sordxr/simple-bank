package account

import (
	"context"
	"fmt"
	"github.com/D1sordxr/packages/postgres/executor"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/models"
)

type DAO struct {
	Executor *executor.Manager
}

func NewDAO(exec *executor.Manager) *DAO {
	return &DAO{Executor: exec}
}

func (d *DAO) GetProjection(ctx context.Context, id string) (model models.Account, err error) {
	const op = "postgres.AccountDAO.GetProjection"

	conn := d.Executor.GetExecutor(ctx)

	query := `
		SELECT id, 
		available_money,
		status
		FROM accounts WHERE id = $1
	`

	err = conn.QueryRow(ctx, query, id).Scan(&model.ID, &model.AvailableMoney, &model.Status)
	if err != nil {
		return models.Account{}, fmt.Errorf("%s: %w: %w", op, ErrReadingAccount, err)
	}

	return model, nil
}

func (d *DAO) UpdateProjection(ctx context.Context, model models.Account) error {
	const op = "postgres.AccountDAO.UpdateProjection"

	conn := d.Executor.GetExecutor(ctx)

	query := `
		UPDATE accounts SET available_money = $1, status = $2 WHERE id = $3
	`

	_, err := conn.Exec(ctx, query, model.AvailableMoney, model.Status, model.ID)
	if err != nil {
		return fmt.Errorf("%s: %w: %w", op, ErrUpdatingAccount, err)
	}

	return nil
}
