package uow

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/persistence"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres"
	"github.com/jackc/pgx/v5"
)

var (
	ctx = context.Background()
)

type UoW struct {
	Conn *postgres.Connection
	Tx   pgx.Tx
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

func (u *UoW) Begin() (interface{}, error) {
	tx, err := u.Conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	u.Tx = tx
	return u.Tx, nil
}

type UoWManager struct {
	Conn *postgres.Connection
}

func (u *UoWManager) GetUoW() persistence.UnitOfWork {
	return &UoW{
		Conn: u.Conn,
	}
}

func NewUoWManager(conn *postgres.Connection) *UoWManager {
	return &UoWManager{
		Conn: conn,
	}
}
