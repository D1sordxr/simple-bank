package interfaces

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/account/commands"
)

type CreateAccountCommand interface {
	Handle(ctx context.Context, c commands.CreateAccountCommand) (commands.CreateDTO, error)
}
