package client

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, client Aggregate) error
	Update(ctx context.Context, client Aggregate) error
	Exists(ctx context.Context, email string) error
}
