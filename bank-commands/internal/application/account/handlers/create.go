package handlers

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-banking-system/internal/application/account/commands"
	accountRoot "github.com/D1sordxr/simple-banking-system/internal/domain/account"
	"github.com/D1sordxr/simple-banking-system/internal/domain/account/vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
	sharedExceptions "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_exceptions"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_vo"
)

type CreateAccountHandler struct {
	deps *commands.Dependencies
}

func NewCreateAccountHandler(dependencies *commands.Dependencies) *CreateAccountHandler {
	return &CreateAccountHandler{
		deps: dependencies,
	}
}

func (h *CreateAccountHandler) Handle(ctx context.Context, c commands.CreateAccountCommand) (commands.CreateDTO, error) {
	const op = "Services.AccountService.CreateAccount"

	logger := h.deps.Logger
	log := logger.With(
		logger.String("operation", op),
		logger.Group("account",
			logger.String("clientID", c.ClientID),
			logger.String("currency", c.Currency),
		),
	)

	log.Info("Attempting to create new account")

	clientID, err := sharedVO.NewUUIDFromString(c.ClientID)
	if err != nil {
		log.Error(sharedExceptions.LogVOCreationError("UUID"), sharedExceptions.LogError(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	accountID := sharedVO.NewUUID()

	balance := vo.NewBalance()

	currency, err := sharedVO.NewCurrency(c.Currency)
	if err != nil {
		log.Error(sharedExceptions.LogVOCreationError("currency"), sharedExceptions.LogError(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	status := vo.NewStatus()

	account, err := accountRoot.NewAccount(clientID, accountID, balance, currency, status)
	if err != nil {
		log.Error(sharedExceptions.LogAggregateCreationError("account"), sharedExceptions.LogError(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	uow := h.deps.UoWManager.GetUoW()
	tx, err := uow.Begin()
	if err != nil {
		log.Error(sharedExceptions.LogErrorAsString(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}
	defer func() {
		if r := recover(); r != nil {
			_ = uow.Rollback()
			panic(r)
		}
		if err != nil {
			log.Error(sharedExceptions.LogErrorAsString(err))
			_ = uow.Rollback()
		}
	}()

	if err = h.deps.AccountRepository.Create(ctx, tx, account); err != nil {
		log.Error(sharedExceptions.LogErrorAsString(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	accountEvent, err := event.NewAccountCreatedEvent(account)
	if err != nil {
		log.Error(sharedExceptions.LogEventCreationError(), sharedExceptions.LogError(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}
	if err = h.deps.EventRepository.SaveEvent(ctx, tx, accountEvent); err != nil {
		log.Error(sharedExceptions.LogErrorAsString(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	outboxEvent, err := outbox.NewOutboxEvent(accountEvent)
	if err != nil {
		log.Error(sharedExceptions.LogOutboxCreationError(), sharedExceptions.LogError(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}
	if err = h.deps.OutboxRepository.SaveOutboxEvent(ctx, tx, outboxEvent); err != nil {
		log.Error(sharedExceptions.LogErrorAsString(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	if err = uow.Commit(); err != nil {
		log.Error(sharedExceptions.LogErrorAsString(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Account creation completed successfully")
	return commands.CreateDTO{
		AccountID: accountID.String(),
	}, nil
}
