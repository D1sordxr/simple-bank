package outbox

import "context"

type Repository interface {
	SaveOutboxEvent(ctx context.Context, tx interface{}, outbox Outbox) error
}
