package context_utils

import (
	"context"
	"github.com/jackc/pgx/v5"
)

const (
	TxCtxKey    = "tx"
	BatchCtxKey = "batch"
)

func GetTxFromContext(ctx context.Context) (pgx.Tx, error) {
	tx := ctx.Value(TxCtxKey)
	if tx == nil {
		return nil, ErrNoTxInCtx
	}
	return tx.(pgx.Tx), nil
}

func GetBatchFromContext(ctx context.Context) (*pgx.Batch, error) {
	batch := ctx.Value(BatchCtxKey)
	if batch == nil {
		return nil, ErrNoBatchInCtx
	}
	return batch.(*pgx.Batch), nil
}
