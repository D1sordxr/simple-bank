package outbox

import "context"

type Repository interface {
	FetchPendingMessages(ctx context.Context, tx interface{}, limit int) ([]Aggregate, error)
	MarkAsProcessed(ctx context.Context, tx interface{}, id string) error
}
