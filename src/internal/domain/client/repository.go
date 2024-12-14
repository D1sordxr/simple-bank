package client

import "context"

type Repository interface {
	Create(ctx context.Context, tx interface{}, client Aggregate) error
	Update(ctx context.Context, tx interface{}, client Aggregate) error
	Load(ctx context.Context, email string) (Aggregate, error)
	Exists(ctx context.Context, email string) error
}
