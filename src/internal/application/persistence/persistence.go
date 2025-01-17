package persistence

import (
	"context"
)

type UoWManager interface {
	GetUoW() UnitOfWork
}

type UnitOfWork interface {
	Begin() (interface{}, error)
	Commit() error
	Rollback() error

	// TODO: Clients() ClientRepository
	// TODO: Accounts() AccountRepository
	// TODO: Transfers() TransferRepository

	// TODO: PublishEvents(ctx context.Context) error
}

type TransactionOutbox interface {
	CreateEvent(ctx context.Context, event []byte) error
}
