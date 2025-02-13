package client

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/client/interfaces"
)

type Service struct {
	interfaces.CreateClientCommand
	interfaces.UpdateClientCommand
}

func NewClientService(
	create interfaces.CreateClientCommand,
	update interfaces.UpdateClientCommand,
) *Service {
	return &Service{
		CreateClientCommand: create,
		UpdateClientCommand: update,
	}
}
