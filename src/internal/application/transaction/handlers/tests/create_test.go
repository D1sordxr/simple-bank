package tests

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/transaction/commands"
	"github.com/D1sordxr/simple-banking-system/internal/application/transaction/handlers"
	"github.com/google/uuid"
	"testing"
)

const (
	transferType   = "transfer"
	depositType    = "deposit"
	withdrawalType = "withdrawal"
	reversalType   = "reversal"
)

// TODO: fix nil uuid vo creation

// TODO: tests with DepositType, WithdrawalType, ReversalType

func TestSuccessCreateTransactionHandlerWithTransferType(t *testing.T) {
	command := commands.CreateTransactionCommand{
		SourceAccountID:      uuid.New().String(),
		DestinationAccountID: uuid.New().String(),
		Currency:             "RUB",
		Amount:               1,
		Type:                 transferType,
		Description:          "something very informative",
	}

	ctx := context.Background()

	mockDeps := mockClientDeps()

	txService := handlers.NewCreateTransactionHandler(mockDeps)

	response, err := txService.Handle(ctx, command)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	t.Logf("TransactionID: %s", response.TransactionID)
}
