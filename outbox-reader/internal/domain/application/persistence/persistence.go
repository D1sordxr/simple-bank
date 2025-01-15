package persistence

import "context"

type UoW interface {
	Commit() error
	Rollback() error
	Begin() (interface{}, error)
}

type UoWManager interface {
	GetUoW() UoW
}

type Producer interface {
	SendMessage(ctx context.Context, key, value []byte) error
	Close() error
}
