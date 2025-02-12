package account

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-banking-system/internal/domain/account"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres"
	converters "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/converters/account"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn *postgres.Connection
}

func NewAccountRepository(conn *postgres.Connection) *Repository {
	return &Repository{Conn: conn}
}

func (r *Repository) Create(ctx context.Context, tx interface{}, account account.Aggregate) error {
	const op = "postgres.AccountRepository.Create"

	conn := tx.(pgx.Tx)

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

func (r *Repository) GetByID(ctx context.Context, accountID uuid.UUID) (account.Aggregate, error) {
	const (
		op  = "postgres.AccountRepository.GetByID"
		msg = "not implemented"
	)

	return account.Aggregate{}, fmt.Errorf("%s: %s", op, msg)
}
