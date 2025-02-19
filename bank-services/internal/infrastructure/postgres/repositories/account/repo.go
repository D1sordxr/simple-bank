package account

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/account"
	converters "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/converters/account"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/executor"
)

type Repository struct {
	Executor *executor.Executor
}

func NewAccountRepository(executor *executor.Executor) *Repository {
	return &Repository{Executor: executor}
}

func (r *Repository) Create(ctx context.Context, account account.Aggregate) error {
	const op = "postgres.AccountRepository.Create"

	conn := r.Executor.GetExecutor(ctx)

	accountsQuery := `INSERT INTO accounts (
        id, 
        client_id, 
        available_money,
        frozen_money,
        currency,
        status,
        created_at,
        updated_at
    ) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	accountModel := converters.ConvertAggregateToModel(account)

	_, err := conn.Exec(ctx, accountsQuery,
		accountModel.ID,
		accountModel.ClientID,
		accountModel.AvailableMoney,
		accountModel.FrozenMoney,
		accountModel.Currency,
		accountModel.Status,
		accountModel.CreatedAt,
		accountModel.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrFailedToCreateAccount)
	}

	return nil
}
