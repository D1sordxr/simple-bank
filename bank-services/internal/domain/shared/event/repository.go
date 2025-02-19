package event

import "context"

type Repository interface {
	SaveEvent(ctx context.Context, event Event) error
}
