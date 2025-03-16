package account

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/interfaces"
)

type Service struct {
	interfaces.CreateAccountCommand
	interfaces.UpdateAccountCommand
}

func NewAccountService(
	create interfaces.CreateAccountCommand,
	update interfaces.UpdateAccountCommand,
) *Service {
	return &Service{
		CreateAccountCommand: create,
		UpdateAccountCommand: update,
	}
}
