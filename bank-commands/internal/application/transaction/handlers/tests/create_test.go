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

func FuzzCreateTransactionHandler(f *testing.F) {
	seedCorpus := []commands.CreateTransactionCommand{
		{ // Transfer
			SourceAccountID:      uuid.NewString(),
			DestinationAccountID: uuid.NewString(),
			Currency:             "RUB",
			Amount:               100,
			Type:                 transferType,
			Description:          "valid transfer",
		},
		{ // Deposit
			DestinationAccountID: uuid.NewString(),
			Currency:             "USD",
			Amount:               50,
			Type:                 depositType,
			Description:          "valid deposit",
		},
		{ // Withdrawal
			SourceAccountID: uuid.NewString(),
			Currency:        "EUR",
			Amount:          75,
			Type:            withdrawalType,
			Description:     "valid withdrawal",
		},
	}

	for _, cmd := range seedCorpus {
		f.Add(
			cmd.SourceAccountID,
			cmd.DestinationAccountID,
			cmd.Currency,
			cmd.Amount,
			cmd.Type,
			cmd.Description,
		)
	}

	f.Fuzz(func(t *testing.T,
		sourceID string,
		destID string,
		currency string,
		amount float64,
		txType string,
		description string,
	) {
		ctx := context.Background()
		mockDeps := mockClientDeps()
		handler := handlers.NewCreateTransactionHandler(mockDeps)

		cmd := commands.CreateTransactionCommand{
			SourceAccountID:      sourceID,
			DestinationAccountID: destID,
			Currency:             currency,
			Amount:               amount,
			Type:                 txType,
			Description:          description,
		}

		response, err := handler.Handle(ctx, cmd)

		switch cmd.Type {
		case transferType:
			if cmd.SourceAccountID == "" || cmd.DestinationAccountID == "" {
				if err == nil {
					t.Error("expected error for missing accounts in transfer")
				}
				return
			}
		case depositType:
			if cmd.DestinationAccountID == "" {
				if err == nil {
					t.Error("expected error for missing destination in deposit")
				}
				return
			}
		case withdrawalType:
			if cmd.SourceAccountID == "" {
				if err == nil {
					t.Error("expected error for missing source in withdrawal")
				}
				return
			}
		default:
			if err == nil {
				t.Error("expected error for unknown transaction type")
			}
			return
		}

		if err == nil {
			if response.TransactionID == "" {
				t.Error("expected non-empty transaction ID")
			}
		} else {
			t.Logf("expected error: %v", err)
		}
	})
}
