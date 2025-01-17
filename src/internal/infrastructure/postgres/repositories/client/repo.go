package client

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres"
)

// TODO: implement methods

type Repository struct {
	Conn *postgres.Connection
}

func NewClientRepository(conn *postgres.Connection) *Repository {
	return &Repository{Conn: conn}
}

func (r *Repository) Create(ctx context.Context, tx interface{}, aggregate client.Aggregate) error {
	return nil
}

func (r *Repository) SaveEvent(ctx context.Context, tx interface{}, aggregate client.Aggregate) error {
	return nil
}

func (r *Repository) SaveOutboxEvent(ctx context.Context, tx interface{}, outbox outbox.Outbox) error {
	return nil
}
