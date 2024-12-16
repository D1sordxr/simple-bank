package account

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, tx interface{}, account Aggregate) error
	GetByID(ctx context.Context, accountID uuid.UUID) (Aggregate, error)
	ClientExists(ctx context.Context, clientID uuid.UUID) error // could be omitted near future
}
