package outbox

import "context"

type Repository interface {
	SaveOutboxEvent(ctx context.Context, outbox Outbox) error
}
