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
	Float64(key string, v float64) slog.Attr
	Group(key string, args ...any) slog.Attr
}

type UoW interface {
	Commit() error
	Rollback() error
	Begin() (interface{}, error)
}

type Producer interface {
	SendMessage(ctx context.Context, key, value []byte) error
	Close() error
}
