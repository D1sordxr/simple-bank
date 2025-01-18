package client

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
)

type Repository interface {
	Create(ctx context.Context, tx interface{}, client Aggregate) error
	Update(ctx context.Context, tx interface{}, client Aggregate) error
	Load(ctx context.Context, email string) (Aggregate, error)
	Exists(ctx context.Context, email string) error

	// SaveEvent saves an event to EventStore
	SaveEvent(ctx context.Context, tx interface{}, event event.Event) error

	// SaveOutboxEvent saves an event to Outbox
	SaveOutboxEvent(ctx context.Context, tx interface{}, outbox outbox.Outbox) error
}
