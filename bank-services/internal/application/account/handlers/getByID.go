package handlers

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/dependencies"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/queries"
)

type GetByIDAccountHandler struct {
	deps *dependencies.Dependencies
}

func NewGetByIDAccountHandler(deps *dependencies.Dependencies) *GetByIDAccountHandler {
	return &GetByIDAccountHandler{
		deps: deps,
	}
}

func (h *GetByIDAccountHandler) Handle(ctx context.Context, q queries.GetByIDAccountQuery) (queries.GetByIDAccountDTO, error) {

	// TODO: ...

	return queries.GetByIDAccountDTO{}, nil
}
