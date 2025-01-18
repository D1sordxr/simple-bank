package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/client/commands"
	"github.com/D1sordxr/simple-banking-system/internal/application/persistence"
	clientRoot "github.com/D1sordxr/simple-banking-system/internal/domain/client"
)

// TODO: UpdateClientHandler

type UpdateClientHandler struct {
	clientRoot.Repository
	persistence.UoWManager
}

func NewUpdateClientHandler(repo clientRoot.Repository,
	uow persistence.UoWManager) *UpdateClientHandler {
	return &UpdateClientHandler{
		UoWManager: uow,
		Repository: repo,
	}
}

func (h *UpdateClientHandler) Handle(ctx context.Context, command commands.CreateClientCommand) (commands.CreateDTO, error) {
	_ = command
	_ = ctx
	return commands.CreateDTO{}, nil
}
