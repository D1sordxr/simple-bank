package interfaces

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/commands"
)

type OutboxCommand interface {
	UpdateStatus(ctx context.Context, command commands.OutboxCommand) error
}
