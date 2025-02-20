package persistence

import (
	"context"
	"log/slog"
)

type Logger interface {
	Info(msg string, attrs ...any)
	Error(msg string, attrs ...any)
	Debug(msg string, attrs ...any)
	String(key string, value string) slog.Attr
	Int(key string, value int) slog.Attr
	Float64(key string, v float64) slog.Attr
	Group(key string, args ...any) slog.Attr
}

type UnitOfWork interface {
	BeginWithTxAndBatch(ctx context.Context) (context.Context, error)
	BeginWithTx(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) error
	GracefulRollback(ctx context.Context, err *error)
	Commit(ctx context.Context) error
}

type Producer interface {
	SendMessage(ctx context.Context, key, value []byte) error
	Close() error
}
