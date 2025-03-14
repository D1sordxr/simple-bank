package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/client"
	converters "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/converters/client"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/executor"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Executor *executor.Executor
}

func NewClientRepository(executor *executor.Executor) *Repository {
	return &Repository{Executor: executor}
}

func (r *Repository) Create(ctx context.Context, client client.Aggregate) error {
	const op = "postgres.ClientRepository.Create"

	conn := r.Executor.GetExecutor(ctx)

	clientsQuery := `INSERT INTO clients (
        id, 
        first_name, 
        last_name, 
        middle_name, 
        email,
        status,
        created_at,
        updated_at
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	phonesQuery := `INSERT INTO phones (
        id,
        client_id,
        phone_number,
        created_at,
        updated_at
    ) VALUES ($1, $2, $3, $4, $5)`

	clientModel := converters.ConvertAggregateToModel(client)

	_, err := conn.Exec(ctx, clientsQuery,
		clientModel.ID,
		clientModel.FirstName,
		clientModel.LastName,
		clientModel.MiddleName,
		clientModel.Email,
		clientModel.Status,
		clientModel.CreatedAt,
		clientModel.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrFailedToCreateClient)
	}

	// TODO: add batch

	for _, phone := range client.Phones {
		phoneModel := converters.ConvertPhoneEntityToModel(phone)

		_, err = conn.Exec(ctx, phonesQuery,
			phoneModel.ID,
			phoneModel.ClientID,
			phoneModel.PhoneNumber,
			phoneModel.CreatedAt,
			phoneModel.UpdatedAt,
		)
		if err != nil {
			return fmt.Errorf("%s: %w", op, ErrFailedToCreatePhone)
		}
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, client client.Aggregate) error {

	// TODO: ...

	return nil
}

func (r *Repository) Exists(ctx context.Context, email string) error {
	const op = "postgres.ClientRepository.Exists"

	conn := r.Executor.GetExecutor(ctx)

	query := `SELECT email FROM clients WHERE email = $1`

	var existingEmail string

	err := conn.QueryRow(ctx, query, email).Scan(&existingEmail)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	if email == existingEmail {
		return fmt.Errorf("%s: %w", op, ErrClientAlreadyExists)
	}
	return nil
}
