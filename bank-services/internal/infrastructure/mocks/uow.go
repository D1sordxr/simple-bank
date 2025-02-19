package mocks

import (
	"context"
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
func (t *TestUoW) BeginSerializableTxWithRetry() (interface{}, error) {
	_ = ctx
	return nil, nil
}
func (t *TestUoW) BeginSerializableTx() (interface{}, error) {
	_ = ctx
	return nil, nil
}
