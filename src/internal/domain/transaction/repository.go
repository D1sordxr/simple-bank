package transaction

import "context"

type Repository interface {
	Create(ctx context.Context, tx interface{}, transaction Aggregate) error

	// TODO: SaveEvent(ctx, tx, event) // Saves event to EventStore

	// TODO: SaveOutboxEvent(ctx, tx, outboxEvent) // Saves to Outbox

}
