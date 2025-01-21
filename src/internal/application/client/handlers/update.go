package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/client"
	"github.com/D1sordxr/simple-banking-system/internal/application/client/commands"
)

// TODO: UpdateClientHandler

type UpdateClientHandler struct {
	*client.Dependencies
}

func NewUpdateClientHandler(dependencies *client.Dependencies) *UpdateClientHandler {
	return &UpdateClientHandler{Dependencies: dependencies}
}

func (h *UpdateClientHandler) Handle(ctx context.Context, command commands.CreateClientCommand) (commands.CreateDTO, error) {
	_ = command
	_ = ctx
	return commands.CreateDTO{}, nil
}
