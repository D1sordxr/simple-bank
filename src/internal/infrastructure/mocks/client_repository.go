package mocks

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client"
	"github.com/stretchr/testify/mock"
)

type MockClientRepository struct {
	mock.Mock
}

func (t *MockClientRepository) Create(ctx context.Context, tx interface{}, client client.Aggregate) error {
	args := t.Called(ctx, tx, client)
	return args.Error(0)
}

func (t *MockClientRepository) Update(ctx context.Context, tx interface{}, client client.Aggregate) error {
	args := t.Called(ctx, tx, client)
	return args.Error(0)
}

func (t *MockClientRepository) Load(ctx context.Context, email string) (client.Aggregate, error) {
	args := t.Called(ctx, email)
	user, ok := args.Get(0).(client.Aggregate)
	if !ok {
		return client.Aggregate{}, args.Error(1)
	}
	return user, args.Error(1)
}

func (t *MockClientRepository) Exists(ctx context.Context, email string) error {
	args := t.Called(ctx, email)
	return args.Error(0)
}
