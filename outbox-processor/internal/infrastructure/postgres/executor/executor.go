package executor

import (
	"context"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres"
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

type (
	txKey    struct{}
	batchKey struct{}
)

type Executor struct {
	*postgres.Pool
}

func NewExecutor(pool *postgres.Pool) *Executor {
	return &Executor{Pool: pool}
}

func (e *Executor) InjectTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func (e *Executor) ExtractTx(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(txKey{}).(pgx.Tx)
	return tx, ok
}

func (e *Executor) NewBatch() *BatchExecutor {
	return &BatchExecutor{Batch: &pgx.Batch{}}
}

func (e *Executor) InjectBatch(ctx context.Context, batch *BatchExecutor) context.Context {
	return context.WithValue(ctx, batchKey{}, batch)
}

func (e *Executor) ExtractBatch(ctx context.Context) (*BatchExecutor, bool) {
	batch, ok := ctx.Value(batchKey{}).(*BatchExecutor)
	return batch, ok
}

func (e *Executor) GetExecutor(ctx context.Context) IExecutor {
	if batch, ok := e.ExtractBatch(ctx); ok {
		return batch
	}

	if tx, ok := e.ExtractTx(ctx); ok {
		return tx
	}

	return e.Pool
}
