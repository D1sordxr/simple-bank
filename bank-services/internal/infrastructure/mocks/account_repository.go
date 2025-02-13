package mocks

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/account"
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

func (t *MockAccountRepository) GetByID(ctx context.Context, clientID uuid.UUID) (account.Aggregate, error) {
	args := t.Called(ctx, clientID)
	accountData, ok := args.Get(0).(account.Aggregate)
	if !ok {
		return account.Aggregate{}, args.Error(1)
	}
	return accountData, args.Error(1)
}

func (t *MockAccountRepository) ClientExists(ctx context.Context, clientID uuid.UUID) error {
	args := t.Called(ctx, clientID)
	return args.Error(0)
}
