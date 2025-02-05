package handlers

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-banking-system/internal/application/client/commands"
	clientRoot "github.com/D1sordxr/simple-banking-system/internal/domain/client"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/entity"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
	sharedExceptions "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_exceptions"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_vo"
)

type CreateClientHandler struct {
	deps *commands.Dependencies
}

func NewCreateClientHandler(dependencies *commands.Dependencies) *CreateClientHandler {
	return &CreateClientHandler{deps: dependencies}
}

func (h *CreateClientHandler) Handle(ctx context.Context, c commands.CreateClientCommand) (commands.CreateDTO, error) {
	const op = "Services.ClientService.CreateClient"

	logger := h.deps.Logger
	log := logger.With(
		logger.String("operation", op),
		logger.String("clientEmail", c.Email),
	)

	log.Info("Attempting to create new client")

	clientID := sharedVO.NewUUID()

	fullName, err := vo.NewFullName(c.FirstName, c.MiddleName, c.LastName)
	if err != nil {
		log.Error(sharedExceptions.LogVOCreationError("fullName"), sharedExceptions.LogError(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	email, err := vo.NewEmail(c.Email)
	if err != nil {
		log.Error(sharedExceptions.LogVOCreationError("email"), sharedExceptions.LogError(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	phones, err := entity.NewPhones(c.Phones, clientID.Value)
	if err != nil {
		log.Error(sharedExceptions.LogEntityCreationError("phones"), sharedExceptions.LogError(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}
	status := vo.NewStatus()

	err = h.deps.ClientRepository.Exists(ctx, email.Email)
	if err != nil {
		log.Error(sharedExceptions.LogErrorAsString(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	client, err := clientRoot.NewClient(clientID, fullName, email, phones, status)
	if err != nil {
		log.Error(sharedExceptions.LogAggregateCreationError("client"))
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

	err = h.deps.ClientRepository.Create(ctx, tx, client)
	if err != nil {
		log.Error(sharedExceptions.LogErrorAsString(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	clientEvent, err := event.NewClientCreatedEvent(client)
	if err != nil {
		log.Error(sharedExceptions.LogEventCreationError(), sharedExceptions.LogError(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}
	if err = h.deps.EventRepository.SaveEvent(ctx, tx, clientEvent); err != nil {
		log.Error(sharedExceptions.LogErrorAsString(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	outboxEvent, err := outbox.NewOutboxEvent(clientEvent)
	if err != nil {
		log.Error(sharedExceptions.LogOutboxCreationError(), sharedExceptions.LogError(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}
	if err = h.deps.OutboxRepository.SaveOutboxEvent(ctx, tx, outboxEvent); err != nil {
		log.Error(sharedExceptions.LogErrorAsString(err))
	}

	if err = uow.Commit(); err != nil {
		log.Error(sharedExceptions.LogErrorAsString(err))
		return commands.CreateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Client creation completed successfully")
	return commands.CreateDTO{
		ClientID: client.ClientID.String(),
		FullName: client.FullName.String(),
		Email:    client.Email.String(),
		Phones:   client.Phones.Read(),
		Status:   client.Status.String(),
	}, nil
}
