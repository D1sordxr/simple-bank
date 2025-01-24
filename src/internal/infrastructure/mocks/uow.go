package mocks

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/persistence"
)

var (
	ctx = context.Background()
)

type TestUoW struct {
}

func (t *TestUoW) Begin() (interface{}, error) {
	_ = ctx
	return nil, nil
}
func (t *TestUoW) Commit() error {
	_ = ctx
	return nil
}
func (t *TestUoW) Rollback() error {
	_ = ctx
	return nil
}

type TestUoWManager struct {
}

func (t *TestUoWManager) GetUoW() persistence.UnitOfWork {
	return &TestUoW{}
}
