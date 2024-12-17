package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/persistence"
	"github.com/D1sordxr/simple-banking-system/internal/application/transaction/commands"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/vo"
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

	transactionID := sharedVO.NewUUID()

	var sourceAccountID, destinationAccountID *sharedVO.UUID
	if len(c.SourceAccountID) != 0 {
		parsedSrcID, err := sharedVO.NewPointerUUIDFromString(c.SourceAccountID)
		if err != nil {
			return commands.CreateTransactionDTO{}, err
		}
		sourceAccountID = parsedSrcID
	}
	if len(c.DestinationAccountID) != 0 {
		parsedDestID, err := sharedVO.NewPointerUUIDFromString(c.DestinationAccountID)
		if err != nil {
			return commands.CreateTransactionDTO{}, err
		}
		destinationAccountID = parsedDestID
	}

	currency, err := sharedVO.NewCurrency(c.Currency)
	if err != nil {
		return commands.CreateTransactionDTO{}, err
	}

	_ = ctx

	return commands.CreateTransactionDTO{}, nil
}
