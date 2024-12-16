package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/account/commands"
	"github.com/D1sordxr/simple-banking-system/internal/application/persistence"
	"github.com/D1sordxr/simple-banking-system/internal/domain/account"
	"github.com/D1sordxr/simple-banking-system/internal/domain/account/exceptions"
	"github.com/google/uuid"
)

type GetByIDAccountHandler struct {
	Repository account.Repository
	UoWManager persistence.UoWManager
}

func NewGetByIDAccountHandler(repo account.Repository,
	uow persistence.UoWManager) *GetByIDAccountHandler {
	return &GetByIDAccountHandler{
		Repository: repo,
		UoWManager: uow,
	}
}

func (h *GetByIDAccountHandler) Handle(ctx context.Context, c commands.GetByIDAccountCommand) (commands.GetByIDAccountDTO, error) {
	accountID, err := uuid.Parse(c.AccountID)
	if err != nil {
		return commands.GetByIDAccountDTO{}, exceptions.InvalidUUID
	}

	accountData, err := h.Repository.GetByID(ctx, accountID)
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
