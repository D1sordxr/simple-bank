package postgres

import (
	"context"
	storageConfig "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres/config"
	"github.com/jackc/pgx/v5"
)

type Connection struct {
	*pgx.Conn
}

func NewConnection(config *storageConfig.StorageConfig) *Connection {
	conn, err := pgx.Connect(context.Background(), config.ConnectionString())
	if err != nil {
		panic(err)
	}
	return &Connection{Conn: conn}
}
