package transaction

import "context"

type Repository interface {
	Create(ctx context.Context, tx interface{}, transaction Aggregate) error
}
