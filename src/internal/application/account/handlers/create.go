package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/account/commands"
	"github.com/D1sordxr/simple-banking-system/internal/application/persistence"
	accountRoot "github.com/D1sordxr/simple-banking-system/internal/domain/account"
	"github.com/D1sordxr/simple-banking-system/internal/domain/account/vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/exceptions"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/vo"
	"github.com/google/uuid"
)

type CreateAccountHandler struct {
	Repository accountRoot.Repository
	UoWManager persistence.UoWManager
}

func NewCreateAccountHandler(repo accountRoot.Repository,
	uow persistence.UoWManager) *CreateAccountHandler {
	return &CreateAccountHandler{
		Repository: repo,
		UoWManager: uow,
	}
}

func (h *CreateAccountHandler) Handle(ctx context.Context, c commands.CreateAccountCommand) (commands.CreateDTO, error) {
	clientID, err := uuid.Parse(c.ClientID)
	if err != nil {
		return commands.CreateDTO{}, err
	}
	accountID := uuid.New()
	balance := vo.NewBalance()
	currency, err := sharedVO.NewCurrency(c.Currency)
	if err != nil {
		return commands.CreateDTO{}, err
	}

	err = h.Repository.ClientExists(ctx, clientID)
	if err != nil {
		return commands.CreateDTO{}, exceptions.ClientIDNotFound
	}

	account, err := accountRoot.NewAccount(clientID, accountID, balance, currency)
	if err != nil {
		return commands.CreateDTO{}, err
	}

	uow := h.UoWManager.GetUoW()
	tx, err := uow.Begin(ctx)
	if err != nil {
		return commands.CreateDTO{}, err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = uow.Rollback(ctx)
			panic(r)
		}
		if err != nil {
			_ = uow.Rollback(ctx)
		}
	}()

	if err = h.Repository.Create(ctx, tx, account); err != nil {
		return commands.CreateDTO{}, err
	}
	if err = uow.Commit(ctx); err != nil {
		return commands.CreateDTO{}, err
	}

	return commands.CreateDTO{
		AccountID: accountID.String(),
	}, nil
}
