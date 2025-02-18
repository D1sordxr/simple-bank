package postgres

import (
	"context"
	storageConfig "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

type Pool struct {
	*pgxpool.Pool
}

func NewPool(config *storageConfig.StorageConfig) *Pool {
	pool, err := pgxpool.New(context.Background(), config.ConnectionString())
	if err != nil {
		panic(err)
	}

	return &Pool{Pool: pool}
}
