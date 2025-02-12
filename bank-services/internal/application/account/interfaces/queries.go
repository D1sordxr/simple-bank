package interfaces

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/account/queries"
)

type GetByIDAccountQuery interface {
	Handle(ctx context.Context, q queries.GetByIDAccountQuery) (queries.GetByIDAccountDTO, error)
}
