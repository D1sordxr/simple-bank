package client

import "context"

type Repository interface {
	Create(client Aggregate, tx interface{}) error
	Update(client Aggregate, tx interface{}) error
	Load(clientID string) (Aggregate, error)
	Exists(ctx context.Context, email string) error
}
