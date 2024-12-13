package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/client/commands"
)

type UpdateClientHandler struct {
}

func (h *UpdateClientHandler) Handle(_ context.Context, command commands.CreateClientCommand) (commands.CreateDTO, error) {
	_ = command
	return commands.CreateDTO{}, nil
}
