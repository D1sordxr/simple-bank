package account

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, account Aggregate) error
}
