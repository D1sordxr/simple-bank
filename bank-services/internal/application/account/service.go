package account

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/interfaces"
)

type Service struct {
	interfaces.CreateAccountCommand
	interfaces.GetByIDAccountQuery
}

func NewAccountService(
	create interfaces.CreateAccountCommand,
	getByID interfaces.GetByIDAccountQuery,
) *Service {
	return &Service{
		CreateAccountCommand: create,
		GetByIDAccountQuery:  getByID,
	}
}
