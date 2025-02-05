package postgres

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/domain/application/persistence"
	"github.com/jackc/pgx/v5"
)

type UoW struct {
	Conn Connection
	Tx   pgx.Tx
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

func (u *UoW) Begin() (interface{}, error) {
	tx, err := u.Conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	u.Tx = tx
	return u.Tx, nil
}

type UoWManager struct {
	Conn Connection
}

func (u *UoWManager) GetUoW() persistence.UoW {
	return &UoW{
		Conn: u.Conn,
	}
}

func NewUoWManager(conn Connection) *UoWManager {
	return &UoWManager{
		Conn: conn,
	}
}
