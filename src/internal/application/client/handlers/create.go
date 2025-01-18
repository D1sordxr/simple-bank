package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/client/commands"
	"github.com/D1sordxr/simple-banking-system/internal/application/persistence"
	clientRoot "github.com/D1sordxr/simple-banking-system/internal/domain/client"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/entity"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/vo"
	"github.com/google/uuid"
)

type CreateClientHandler struct {
	Repository clientRoot.Repository
	UoWManager persistence.UoWManager
}

func NewCreateClientHandler(repo clientRoot.Repository,
	uow persistence.UoWManager) *CreateClientHandler {
	return &CreateClientHandler{
		UoWManager: uow,
		Repository: repo,
	}
}

func (h *CreateClientHandler) Handle(ctx context.Context, c commands.CreateClientCommand) (commands.CreateDTO, error) {
	clientID := uuid.New()
	fullName, err := vo.NewFullName(c.FirstName, c.MiddleName, c.LastName)
	if err != nil {
		return commands.CreateDTO{}, err
	}
	email, err := vo.NewEmail(c.Email)
	if err != nil {
		return commands.CreateDTO{}, err
	}
	phones, err := entity.NewPhones(c.Phones, clientID)
	if err != nil {
		return commands.CreateDTO{}, err
	}
	status := vo.NewStatus()

	err = h.Repository.Exists(ctx, email.Email)
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

	err = h.Repository.Create(ctx, tx, client)
	if err != nil {
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
