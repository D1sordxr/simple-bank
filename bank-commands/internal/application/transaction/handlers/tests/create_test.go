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

// TODO: implement ReversalType -> tests with ReversalType

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

func TestSuccessCreateTransactionHandlerWithDepositType(t *testing.T) {
	command := commands.CreateTransactionCommand{
		DestinationAccountID: uuid.New().String(),
		Currency:             "RUB",
		Amount:               1,
		Type:                 depositType,
		Description:          "deposit transaction",
	}

	ctx := context.Background()
	mockDeps := mockClientDeps()
	txService := handlers.NewCreateTransactionHandler(mockDeps)

	response, err := txService.Handle(ctx, command)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	if response.TransactionID == "" {
		t.Error("expected non-empty transaction ID")
	}
}

func TestSuccessCreateTransactionHandlerWithWithdrawalType(t *testing.T) {
	command := commands.CreateTransactionCommand{
		SourceAccountID: uuid.New().String(),
		Currency:        "RUB",
		Amount:          1,
		Type:            withdrawalType,
		Description:     "withdrawal transaction",
	}

	ctx := context.Background()
	mockDeps := mockClientDeps()
	txService := handlers.NewCreateTransactionHandler(mockDeps)

	response, err := txService.Handle(ctx, command)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	if response.TransactionID == "" {
		t.Error("expected non-empty transaction ID")
	}
}
