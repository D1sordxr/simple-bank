package interfaces

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/outbox"
)

type EventRepo interface {
	SaveEvent(ctx context.Context, event event.Event) error
}

type OutboxRepo interface {
	SaveOutboxEvent(ctx context.Context, outbox outbox.Outbox) error
}
