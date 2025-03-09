package handlers

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/commands"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/dependencies"
)

type UpdateAccountHandler struct {
	deps *dependencies.Dependencies
}

func NewUpdateAccountHandler(dependencies *dependencies.Dependencies) *UpdateAccountHandler {
	return &UpdateAccountHandler{
		deps: dependencies,
	}
}

func (h *UpdateAccountHandler) Handle(ctx context.Context, c commands.UpdateAccountCommand) (commands.CreateDTO, error) {
	const op = "Services.AccountService.CreateAccount"
	return commands.CreateDTO{}, nil
}
