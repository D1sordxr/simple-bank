package account

import (
	"github.com/D1sordxr/simple-banking-system/internal/application/account/handlers"
)

type Service struct {
	*handlers.CreateAccountHandler
	*handlers.GetByIDAccountHandler
}

func NewAccountService(
	create *handlers.CreateAccountHandler,
	getByID *handlers.GetByIDAccountHandler,
) *Service {
	return &Service{
		CreateAccountHandler:  create,
		GetByIDAccountHandler: getByID,
	}
}
