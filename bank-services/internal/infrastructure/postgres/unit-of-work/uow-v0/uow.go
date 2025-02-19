package uow_v0

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	maxSerializeRetries = 5
	ctx                 = context.Background()
)

type UoW struct {
	Conn *postgres.Pool
	Tx   pgx.Tx
}

func NewUoW(pool *postgres.Pool) *UoW {
	return &UoW{Conn: pool}
}

func (u *UoW) Begin() (interface{}, error) {
	tx, err := u.Conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	u.Tx = tx

	return u.Tx, nil
}

func (u *UoW) BeginSerializableTx() (interface{}, error) {
	tx, err := u.Conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return nil, err
	}
	u.Tx = tx

	return u.Tx, nil
}

func (u *UoW) Commit() error {
	return u.Tx.Commit(ctx)
}

func (u *UoW) Rollback() error {
	err := u.Tx.Rollback(ctx)
	if err != nil {
		return err
	}
	return nil
}

// BeginSerializableTxWithRetry if serialize failure happens
func (u *UoW) BeginSerializableTxWithRetry() (interface{}, error) {
	var (
		tx  pgx.Tx
		err error
	)

	for i := 0; i < maxSerializeRetries; i++ {
		tx, err = u.Conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
		if err == nil {
			return tx, nil
		}
		if !isSerializationFailure(err) {
			return nil, err
		}
	}
	return nil, err
}

func isSerializationFailure(err error) bool {
	if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "40001" {
		return true
	}
	return false
}
