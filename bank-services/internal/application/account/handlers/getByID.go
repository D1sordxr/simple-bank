package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/account/dependencies"
	"github.com/D1sordxr/simple-banking-system/internal/application/account/queries"
	"github.com/D1sordxr/simple-banking-system/internal/domain/account/exceptions"
	"github.com/google/uuid"
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

	accountID, err := uuid.Parse(q.AccountID)
	if err != nil {
		return queries.GetByIDAccountDTO{}, exceptions.InvalidUUID
	}

	accountData, err := h.deps.AccountRepository.GetByID(ctx, accountID)
	if err != nil {
		return queries.GetByIDAccountDTO{}, err
	}

	return queries.GetByIDAccountDTO{
		AvailableMoney: accountData.Balance.AvailableMoney.Value,
		FrozenMoney:    accountData.Balance.FrozenMoney.Value,
		Currency:       accountData.Currency.String(),
		Status:         accountData.Status.String(),
		CreatedAt:      accountData.CreatedAt.String(),
		UpdatedAt:      accountData.UpdatedAt.String(),
	}, nil
}
