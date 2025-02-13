package mocks

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (t *MockTransactionRepository) Create(ctx context.Context, tx interface{}, transaction transaction.Aggregate) error {
	args := t.Called(ctx, tx, transaction)
	return args.Error(0)
}
