package persistence

import (
	"context"
)

type TransactionService interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Begin(ctx context.Context) (interface{}, error)
}

type TransactionManager interface {
	GetTxManager() TransactionService
}
