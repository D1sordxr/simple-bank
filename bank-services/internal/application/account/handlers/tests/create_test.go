package tests

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/commands"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/handlers"
	"github.com/google/uuid"
	"testing"
)

func TestSuccessCreateAccountHandler(t *testing.T) {
	command := commands.CreateAccountCommand{
		ClientID: uuid.New().String(),
		Currency: "RUB",
	}

	ctx := context.Background()
	mockDeps := mockAccountDeps()

	createAccount := handlers.NewCreateAccountHandler(mockDeps)

	response, err := createAccount.Handle(ctx, command)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	t.Logf("\n"+
		"AccountID: %s",
		response.AccountID)
}
