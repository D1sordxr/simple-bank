package mocks

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
	"github.com/stretchr/testify/mock"
)

type MockOutboxRepository struct {
	mock.Mock
}

func (t *MockOutboxRepository) SaveOutboxEvent(ctx context.Context, tx interface{}, outbox outbox.Outbox) error {
	args := t.Called(ctx, tx, outbox)
	return args.Error(0)
}
