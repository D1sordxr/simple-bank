package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type UoW struct {
	Conn *Connection
	Tx   pgx.Tx
}

func (u *UoW) Begin() (interface{}, error) {
	tx, err := u.Conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	u.Tx = tx
	return u.Tx, nil
}

func (u *UoW) Commit() error {
	return u.Tx.Commit(context.Background())
}

func (u *UoW) Rollback() error {
	err := u.Tx.Rollback(context.Background())
	if err != nil {
		return err
	}
	return nil
}
