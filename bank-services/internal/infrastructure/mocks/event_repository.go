package mocks

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event"
	"github.com/stretchr/testify/mock"
)

type MockEventRepository struct {
	mock.Mock
}

func (t *MockEventRepository) SaveEvent(ctx context.Context, tx interface{}, event event.Event) error {
	args := t.Called(ctx, tx, event)
	return args.Error(0)
}
