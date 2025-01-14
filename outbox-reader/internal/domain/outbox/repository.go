package outbox

import "context"

type Repository interface {
	FetchPendingMessages(ctx context.Context, limit int) ([]Outbox, error)
	MarkAsProcessed(ctx context.Context, id string) error
}
