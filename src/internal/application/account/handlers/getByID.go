package handlers

import (
	"context"
	accountDeps "github.com/D1sordxr/simple-banking-system/internal/application/account"
	"github.com/D1sordxr/simple-banking-system/internal/application/account/commands"
	"github.com/D1sordxr/simple-banking-system/internal/domain/account/exceptions"
	"github.com/google/uuid"
)

type GetByIDAccountHandler struct {
	deps *accountDeps.Dependencies
}

func NewGetByIDAccountHandler(deps *accountDeps.Dependencies) *GetByIDAccountHandler {
	return &GetByIDAccountHandler{
		deps: deps,
	}
}

func (h *GetByIDAccountHandler) Handle(ctx context.Context, c commands.GetByIDAccountCommand) (commands.GetByIDAccountDTO, error) {

	// TODO: ...

	accountID, err := uuid.Parse(c.AccountID)
	if err != nil {
		return commands.GetByIDAccountDTO{}, exceptions.InvalidUUID
	}

	accountData, err := h.deps.AccountRepository.GetByID(ctx, accountID)
	if err != nil {
		return commands.GetByIDAccountDTO{}, err
	}

	return commands.GetByIDAccountDTO{
		AvailableMoney: accountData.Balance.AvailableMoney.Value,
		FrozenMoney:    accountData.Balance.FrozenMoney.Value,
		Currency:       accountData.Currency.String(),
		Status:         accountData.Status.String(),
		CreatedAt:      accountData.CreatedAt.String(),
		UpdatedAt:      accountData.UpdatedAt.String(),
	}, nil
}
