package interfaces

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/queries"
)

type OutboxQuery interface {
	FetchMessages(ctx context.Context, query queries.OutboxQuery) (queries.OutboxDTOs, error)
}
