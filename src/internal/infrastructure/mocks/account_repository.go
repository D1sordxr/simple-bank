package mocks

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/domain/account"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func (t *MockAccountRepository) Create(ctx context.Context, tx interface{}, account account.Aggregate) error {
	args := t.Called(ctx, tx, account)
	return args.Error(0)
}

func (t *MockAccountRepository) ClientExists(ctx context.Context, clientID uuid.UUID) error {
	args := t.Called(ctx, clientID)
	return args.Error(0)
}
