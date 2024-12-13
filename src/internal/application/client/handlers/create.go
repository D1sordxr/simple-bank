package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/client/commands"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/entity"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/vo"
)

type CreateClientHandler struct {
}

func NewCreateClientHandler() *CreateClientHandler {
	return &CreateClientHandler{}
}

func (h *CreateClientHandler) Handle(_ context.Context, c commands.CreateClientCommand) (commands.CreateDTO, error) {
	fullName, err := vo.NewFullName(c.FirstName, c.MiddleName, c.LastName)
	if err != nil {
		return commands.CreateDTO{}, err
	}
	email, err := vo.NewEmail(c.Email)
	if err != nil {
		return commands.CreateDTO{}, err
	}
	phones, err := entity.NewPhones(c.Phones)
	if err != nil {
		return commands.CreateDTO{}, err
	}
	status := vo.NewStatus()

	return commands.CreateDTO{
		FullName: fullName.String(),
		Email:    email.String(),
		Phones:   phones.Read(),
		Status:   status.String(),
	}, nil
}
