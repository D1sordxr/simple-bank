package client

import (
	"github.com/D1sordxr/simple-banking-system/internal/application/client/handlers"
)

type Service struct {
	*handlers.CreateClientHandler
	*handlers.UpdateClientHandler
}

func NewClientService(
	create *handlers.CreateClientHandler,
	update *handlers.UpdateClientHandler,
) *Service {
	return &Service{
		CreateClientHandler: create,
		UpdateClientHandler: update,
	}
}
