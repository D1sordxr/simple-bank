package interfaces

import (
	"context"
	"github.com/D1sordxr/packages/postgres/uow"
)

type UnitOfWork interface {
	BeginWithTxAndBatch(ctx context.Context) (context.Context, error)
	BeginWithTx(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) error
	GracefulRollback(ctx context.Context, err *error)
	Commit(ctx context.Context) error
}

type IUnitOfWork interface {
	uow.UnitOfWork
}
