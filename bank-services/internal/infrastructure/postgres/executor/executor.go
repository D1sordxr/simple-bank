package executor

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type IExecutor interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...any) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

type Executor struct {
	*postgres.Pool
}

type txKey struct{}

func (e *Executor) InjectTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func (e *Executor) ExtractTx(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(txKey{}).(pgx.Tx)
	return tx, ok
}

func NewExecutor(pool *postgres.Pool) *Executor {
	return &Executor{Pool: pool}
}

func (e *Executor) GetExecutor(ctx context.Context) IExecutor {
	if tx, ok := e.ExtractTx(ctx); ok {
		return tx
	}
	return e.Pool
}
