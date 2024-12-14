package persistence

import (
	"context"
)

type UoWManager interface {
	GetUoW() UnitOfWork
}

type UnitOfWork interface {
	Begin(ctx context.Context) (interface{}, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error

	// TODO: Clients() ClientRepository
	// TODO: Accounts() AccountRepository
	// TODO: Transfers() TransferRepository

	// TODO: PublishEvents(ctx context.Context) error
}
