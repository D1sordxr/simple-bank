package mocks

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/persistence"
)

type TestUoW struct {
}

func (t *TestUoW) Begin(ctx context.Context) (interface{}, error) {
	_ = ctx
	return nil, nil
}
func (t *TestUoW) Commit(ctx context.Context) error {
	_ = ctx
	return nil
}
func (t *TestUoW) Rollback(ctx context.Context) error {
	_ = ctx
	return nil
}

type TestUoWManager struct {
}

func (t *TestUoWManager) GetUoW() persistence.UnitOfWork {
	return &TestUoW{}
}
