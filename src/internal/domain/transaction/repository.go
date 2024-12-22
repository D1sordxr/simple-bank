package transaction

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
)

type Repository interface {
	Create(ctx context.Context, tx interface{}, transaction Aggregate) error

	// SaveEvent saves an event to EventStore
	SaveEvent(ctx context.Context, tx interface{}, event event.Event) error

	// SaveOutboxEvent saves an event to Outbox
	SaveOutboxEvent(ctx context.Context, tx interface{}, outbox outbox.Outbox) error
}
