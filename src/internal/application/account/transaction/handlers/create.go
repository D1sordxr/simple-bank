package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/account/transaction/commands"
	"github.com/D1sordxr/simple-banking-system/internal/application/persistence"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction"
)

type CreateTransactionHandler struct {
	Repository transaction.Repository
	UoWManager persistence.UoWManager
}

func NewCreateTransactionHandler(repo transaction.Repository,
	uow persistence.UoWManager) *CreateTransactionHandler {
	return &CreateTransactionHandler{
		Repository: repo,
		UoWManager: uow,
	}
}

func (h CreateTransactionHandler) Handle(ctx context.Context,
	c commands.CreateTransactionCommand) (commands.CreateTransactionDTO, error) {

	_ = ctx
	_ = c

	return commands.CreateTransactionDTO{}, nil
}
