package interfaces

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/commands"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
)

type CreateTransactionCommand interface {
	Handle(ctx context.Context, c commands.CreateTransactionCommand) (commands.CreateTransactionDTO, error)
}

type UpdateTransactionCommand interface {
	Handle(ctx context.Context, c commands.UpdateTransactionCommand) (dto.UpdateDTO, error)
}
