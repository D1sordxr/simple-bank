package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres"
	converters "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/converters/client"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn *postgres.Connection
}

func NewClientRepository(conn *postgres.Connection) *Repository {
	return &Repository{Conn: conn}
}

func (r *Repository) Create(ctx context.Context, tx interface{}, client client.Aggregate) error {
	const op = "postgres.ClientRepository.Create"

	conn := tx.(pgx.Tx)

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

func (r *Repository) Update(ctx context.Context, tx interface{}, client client.Aggregate) error {

	// TODO: ...

	return nil
}

func (r *Repository) Load(ctx context.Context, email string) (client.Aggregate, error) {

	// TODO: ...

	return client.Aggregate{}, nil
}

func (r *Repository) Exists(ctx context.Context, email string) error {
	const op = "postgres.ClientRepository.Exists"

	query := `SELECT email FROM clients WHERE email = $1`

	var existingEmail string

	err := r.Conn.QueryRow(ctx, query, email).Scan(&existingEmail)
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
