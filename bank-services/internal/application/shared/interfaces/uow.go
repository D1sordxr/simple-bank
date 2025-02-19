package interfaces

import "context"

type UnitOfWork interface {
	BeginWithTxBatch(ctx context.Context) (context.Context, error)
	BeginWithTx(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}
