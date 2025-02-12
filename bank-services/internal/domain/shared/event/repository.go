package event

import "context"

type Repository interface {
	SaveEvent(ctx context.Context, tx interface{}, event Event) error
}
