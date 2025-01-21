package client

import (
	"context"
	"errors"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres"
	converters "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/converters/client"
	"github.com/jackc/pgx/v5"
)

// TODO: implement methods

type Repository struct {
	Conn *postgres.Connection
}

func NewClientRepository(conn *postgres.Connection) *Repository {
	return &Repository{Conn: conn}
}

func (r *Repository) Create(ctx context.Context, tx interface{}, client client.Aggregate) error {
	conn := tx.(pgx.Tx)
	clientsQuery := `INSERT INTO clients (
                    id, 
                    full_name, 
                    email,
                    status,
                    created_at
                 ) 
                 VALUES ($1, $2, $3, $4, $5);`
	phonesQuery := `INSERT INTO phones (
                    id,
                    client_id,
                    phone_number,
                    country,
                    code,
                    number,
                    created_at
                )
                VALUES ($1, $2, $3, $4, $5, $6, $7);`

	clientModel := converters.ConvertAggregateToModel(client)
	_, err := conn.Exec(ctx, clientsQuery,
		clientModel.ID,
		clientModel.FullName,
		clientModel.Email,
		clientModel.Status,
		clientModel.CreatedAt,
	)
	if err != nil {
		return err
	}

	for _, phone := range client.Phones {
		phoneModel := converters.ConvertPhoneEntityToModel(phone)
		_, err = conn.Exec(ctx, phonesQuery,
			phoneModel.ID,
			phoneModel.ClientID,
			phoneModel.PhoneNumber,
			phoneModel.Country,
			phoneModel.Code,
			phoneModel.Number,
			phoneModel.CreatedAt,
		)
		if err != nil {
			return err
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
	query := `SELECT email FROM clients WHERE email = $1`

	var existingEmail string

	err := r.Conn.QueryRow(ctx, query, email).Scan(&existingEmail)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	if email == existingEmail {
		return ErrClientAlreadyExists
	}
	return nil
}
