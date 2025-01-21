package handlers

import (
	"context"
	clientDependencies "github.com/D1sordxr/simple-banking-system/internal/application/client"
	"github.com/D1sordxr/simple-banking-system/internal/application/client/commands"
	clientRoot "github.com/D1sordxr/simple-banking-system/internal/domain/client"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/entity"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_vo"
)

type CreateClientHandler struct {
	*clientDependencies.Dependencies
}

func NewCreateClientHandler(dependencies *clientDependencies.Dependencies) *CreateClientHandler {
	return &CreateClientHandler{Dependencies: dependencies}
}

func (h *CreateClientHandler) Handle(ctx context.Context, c commands.CreateClientCommand) (commands.CreateDTO, error) {
	clientID := sharedVO.NewUUID()
	fullName, err := vo.NewFullName(c.FirstName, c.MiddleName, c.LastName)
	if err != nil {
		return commands.CreateDTO{}, err
	}
	email, err := vo.NewEmail(c.Email)
	if err != nil {
		return commands.CreateDTO{}, err
	}
	phones, err := entity.NewPhones(c.Phones, clientID.Value)
	if err != nil {
		return commands.CreateDTO{}, err
	}
	status := vo.NewStatus()

	err = h.Dependencies.ClientRepository.Exists(ctx, email.Email)
	if err != nil {
		return commands.CreateDTO{}, err
	}

	client, err := clientRoot.NewClient(clientID, fullName, email, phones, status)
	if err != nil {
		return commands.CreateDTO{}, err
	}

	uow := h.UoWManager.GetUoW()
	tx, err := uow.Begin()
	if err != nil {
		return commands.CreateDTO{}, err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = uow.Rollback()
			panic(r)
		}
		if err != nil {
			_ = uow.Rollback()
		}
	}()

	err = h.Dependencies.ClientRepository.Create(ctx, tx, client)
	if err != nil {
		return commands.CreateDTO{}, err
	}

	clientEvent, err := event.NewClientCreatedEvent(client)
	if err != nil {
		return commands.CreateDTO{}, err
	}
	if err = h.Dependencies.EventRepository.SaveEvent(ctx, tx, clientEvent); err != nil {
		return commands.CreateDTO{}, err
	}

	if err = uow.Commit(); err != nil {
		return commands.CreateDTO{}, err
	}

	return commands.CreateDTO{
		ClientID: clientID.String(),
		FullName: fullName.String(),
		Email:    email.String(),
		Phones:   phones.Read(),
		Status:   status.String(),
	}, nil
}
