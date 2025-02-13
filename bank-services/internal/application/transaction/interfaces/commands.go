package interfaces

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/commands"
)

type CreateTransactionCommand interface {
	Handle(ctx context.Context, c commands.CreateTransactionCommand) (commands.CreateTransactionDTO, error)
}
