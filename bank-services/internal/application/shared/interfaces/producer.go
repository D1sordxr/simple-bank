package interfaces

import "context"

type Producer interface {
	SendMessage(ctx context.Context, key, value []byte) error
	Close() error
}
