package executor

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type BatchExecutor struct {
	Batch *pgx.Batch
}

func (b *BatchExecutor) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	b.Batch.Queue(sql, arguments...)
	return pgconn.CommandTag{}, nil
}

func (b *BatchExecutor) Query(ctx context.Context, sql string, optionsAndArgs ...any) (pgx.Rows, error) {
	b.Batch.Queue(sql, optionsAndArgs...)
	return nil, nil
}

func (b *BatchExecutor) QueryRow(ctx context.Context, sql string, optionsAndArgs ...any) pgx.Row {
	b.Batch.Queue(sql, optionsAndArgs...)
	return nil
}

func (b *BatchExecutor) SendBatch(ctx context.Context, batch *pgx.Batch) pgx.BatchResults {
	return nil
}

func (b *BatchExecutor) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	err := errors.New("not supported")
	return 0, err
}
