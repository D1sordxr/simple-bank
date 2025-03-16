package transaction

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/interfaces"
)

type Service struct {
	interfaces.CreateTransactionCommand
	interfaces.UpdateTransactionCommand
}

func NewTransactionService(
	create interfaces.CreateTransactionCommand,
	update interfaces.UpdateTransactionCommand,
) *Service {
	return &Service{
		CreateTransactionCommand: create,
		UpdateTransactionCommand: update,
	}
}
