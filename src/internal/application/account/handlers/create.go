package handlers

import (
	"context"
	accountDeps "github.com/D1sordxr/simple-banking-system/internal/application/account"
	"github.com/D1sordxr/simple-banking-system/internal/application/account/commands"
	accountRoot "github.com/D1sordxr/simple-banking-system/internal/domain/account"
	"github.com/D1sordxr/simple-banking-system/internal/domain/account/vo"
	sharedExceptions "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_exceptions"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_vo"
	"log/slog"
)

type CreateAccountHandler struct {
	deps *accountDeps.Dependencies
}

func NewCreateAccountHandler(dependencies *accountDeps.Dependencies) *CreateAccountHandler {
	return &CreateAccountHandler{
		deps: dependencies,
	}
}

func (h *CreateAccountHandler) Handle(ctx context.Context, c commands.CreateAccountCommand) (commands.CreateDTO, error) {
	const op = "Services.AccountService.CreateAccount"

	log := h.deps.Logger.With(
		slog.String("operation", op),
		slog.String("clientID", c.ClientID),
	)

	log.Info("Attempting to create new account")

	clientID, err := sharedVO.NewUUIDFromString(c.ClientID)
	if err != nil {
		log.Error(sharedExceptions.LogVOCreationError("UUID"), sharedExceptions.LogError(err))
		return commands.CreateDTO{}, err
	}

	accountID := sharedVO.NewUUID()

	balance := vo.NewBalance()

	currency, err := sharedVO.NewCurrency(c.Currency)
	if err != nil {
		log.Error(sharedExceptions.LogVOCreationError("currency"), sharedExceptions.LogError(err))
		return commands.CreateDTO{}, err
	}

	err = h.Repository.ClientExists(ctx, clientID) // remove
	if err != nil {
		return commands.CreateDTO{}, shared_exceptions.ClientIDNotFound
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
