package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/client/commands"
	"github.com/D1sordxr/simple-banking-system/internal/application/client/dependencies"
)

// TODO: UpdateClientHandler

type UpdateClientHandler struct {
	*dependencies.Dependencies
}

func NewUpdateClientHandler(dependencies *dependencies.Dependencies) *UpdateClientHandler {
	return &UpdateClientHandler{Dependencies: dependencies}
}

func (h *UpdateClientHandler) Handle(ctx context.Context, command commands.CreateClientCommand) (commands.CreateDTO, error) {
	_ = command
	_ = ctx
	return commands.CreateDTO{}, nil
}
