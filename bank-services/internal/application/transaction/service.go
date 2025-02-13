package transaction

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/interfaces"
)

type Service struct {
	interfaces.CreateTransactionCommand
}

func NewTransactionService(
	create interfaces.CreateTransactionCommand,
) *Service {
	return &Service{
		CreateTransactionCommand: create,
	}
}
