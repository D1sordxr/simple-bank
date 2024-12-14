package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/client/commands"
)

// TODO: UpdateClientHandler

type UpdateClientHandler struct {
}

func (h *UpdateClientHandler) Handle(ctx context.Context, command commands.CreateClientCommand) (commands.CreateDTO, error) {
	_ = command
	_ = ctx
	return commands.CreateDTO{}, nil
}
