package handlers

import (
	"LearningArch/internal/application/client/commands"
	"context"
)

type UpdateClientHandler struct {
}

func (h *UpdateClientHandler) Handle(_ context.Context, command commands.CreateClientCommand) (commands.CreateDTO, error) {
	_ = command
	return commands.CreateDTO{}, nil
}
