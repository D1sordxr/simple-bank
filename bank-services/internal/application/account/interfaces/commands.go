package interfaces

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/commands"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/dto"
)

type CreateAccountCommand interface {
	Handle(ctx context.Context, c commands.CreateAccountCommand) (commands.CreateDTO, error)
}

type UpdateAccountCommand interface {
	Handle(ctx context.Context, c commands.UpdateAccountCommand) (dto.UpdateDTO, error)
}
