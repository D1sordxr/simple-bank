package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/D1sordxr/simple-banking-system/internal/application/transaction/commands"
	"github.com/D1sordxr/simple-banking-system/internal/application/transaction/dependencies"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
	sharedExc "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_exceptions"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_vo"
	transactionRoot "github.com/D1sordxr/simple-banking-system/internal/domain/transaction"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction/exceptions"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction/vo"
)

type CreateTransactionHandler struct {
	deps *dependencies.Dependencies
}

func NewCreateTransactionHandler(deps *dependencies.Dependencies) *CreateTransactionHandler {
	return &CreateTransactionHandler{deps: deps}
}

func (h *CreateTransactionHandler) Handle(ctx context.Context,
	c commands.CreateTransactionCommand) (commands.CreateTransactionDTO, error) {
	const op = "Services.TransactionService.CreateTransaction"

	logger := h.deps.Logger
	log := logger.With(
		logger.String("operation", op),
		logger.Group("transaction",
			logger.String("type", c.Type),
			logger.String("sourceAccountID", c.SourceAccountID),
			logger.String("destinationAccountID", c.DestinationAccountID),
			logger.String("currency", c.Currency),
			logger.Float64("amount", c.Amount),
		),
	)

	log.Info("Attempting to create transaction")

	txID := sharedVO.NewUUID()

	txType, err := vo.NewType(c.Type)
	if err != nil {
		log.Error(sharedExc.LogVOCreationError("transactionType"), sharedExc.LogError(err))
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	var sourceAccountID, destinationAccountID *sharedVO.UUID

	switch txType.String() {
	case vo.TransferType:
		sourceAccountID, err = sharedVO.NewPointerUUIDFromString(c.SourceAccountID)
		if err != nil {
			log.Error(sharedExc.LogVOCreationError("sourceAccountUUID"), sharedExc.LogError(err))
			return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
		}
		destinationAccountID, err = sharedVO.NewPointerUUIDFromString(c.DestinationAccountID)
		if err != nil {
			log.Error(sharedExc.LogVOCreationError("destinationAccountUUID"), sharedExc.LogError(err))
			return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
		}

	case vo.DepositType:
		destinationAccountID, err = sharedVO.NewPointerUUIDFromString(c.DestinationAccountID)
		if err != nil {
			log.Error(sharedExc.LogVOCreationError("destinationAccountUUID"), sharedExc.LogError(err))
			return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
		}

	case vo.WithdrawalType:
		sourceAccountID, err = sharedVO.NewPointerUUIDFromString(c.SourceAccountID)
		if err != nil {
			log.Error(sharedExc.LogVOCreationError("sourceAccountUUID"), sharedExc.LogError(err))
			return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
		}

	case vo.ReversalType:
		// TODO: add reversal transaction type support
		err = errors.New("reversal transaction is not implemented")
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)

	default:
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, exceptions.InvalidTxType)
	}

	currency, err := sharedVO.NewCurrency(c.Currency)
	if err != nil {
		log.Error(sharedExc.LogVOCreationError("currency"), sharedExc.LogError(err))
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	money, err := sharedVO.NewMoneyFromFloat(c.Amount)
	if err != nil {
		log.Error(sharedExc.LogVOCreationError("money"), sharedExc.LogError(err))
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	txStatus := vo.NewTransactionStatus()

	description, err := vo.NewDescription(c.Description)
	if err != nil {
		log.Error(sharedExc.LogVOCreationError("description"), sharedExc.LogError(err))
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	transaction, err := transactionRoot.NewTransaction(
		txID, sourceAccountID, destinationAccountID, currency, money, txStatus, txType, description,
	)
	if err != nil {
		log.Error(sharedExc.LogAggregateCreationError("transaction"), sharedExc.LogError(err))
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	uow := h.deps.UoWManager.GetUoW()
	tx, err := uow.BeginSerializableTx()
	if err != nil {
		log.Error(sharedExc.LogErrorAsString(err))
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
	}
	defer func() {
		if r := recover(); r != nil {
			_ = uow.Rollback()
			panic(r)
		}
		if err != nil {
			log.Error(sharedExc.LogErrorAsString(err))
			_ = uow.Rollback()
		}
	}()

	if err = h.deps.TransactionRepository.Create(ctx, tx, transaction); err != nil {
		sharedExc.LogErrorAsString(err)
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	txEvent, err := event.NewTransactionCreatedEvent(transaction)
	if err != nil {
		log.Error(sharedExc.LogEventCreationError(), sharedExc.LogError(err))
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
	}
	if err = h.deps.EventRepository.SaveEvent(ctx, tx, txEvent); err != nil {
		log.Error(sharedExc.LogErrorAsString(err))
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	outboxEvent, err := outbox.NewOutboxEvent(txEvent)
	if err != nil {
		log.Error(sharedExc.LogOutboxCreationError(), sharedExc.LogError(err))
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
	}
	if err = h.deps.OutboxRepository.SaveOutboxEvent(ctx, tx, outboxEvent); err != nil {
		log.Error(sharedExc.LogErrorAsString(err))
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	if err = uow.Commit(); err != nil {
		log.Error(sharedExc.LogErrorAsString(err))
		return commands.CreateTransactionDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Transaction creation completed successfully")
	return commands.CreateTransactionDTO{
		TransactionID: txID.String(),
	}, nil
}
