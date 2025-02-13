package interfaces

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/queries"
)

type GetByIDAccountQuery interface {
	Handle(ctx context.Context, q queries.GetByIDAccountQuery) (queries.GetByIDAccountDTO, error)
}
