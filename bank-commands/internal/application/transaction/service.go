package transaction

import (
	"github.com/D1sordxr/simple-banking-system/internal/application/transaction/handlers"
)

type Service struct {
	*handlers.CreateTransactionHandler
}

func NewTransactionService(
	create *handlers.CreateTransactionHandler,
) *Service {
	return &Service{
		CreateTransactionHandler: create,
	}
}
