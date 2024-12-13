package handlers

import (
	"LearningArch/internal/application/client/commands"
	"LearningArch/internal/domain/client/entity"
	"LearningArch/internal/domain/client/vo"
	"context"
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

	return commands.CreateDTO{
		FullName: fullName.String(),
		Email:    email.String(),
		Phones:   phones.Read(),
	}, nil
}
