package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/account/commands"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSuccessCreateAccountHandler(t *testing.T) {
	command := commands.CreateAccountCommand{
		ClientID: uuid.New().String(),
		Currency: "RUB",
	}

	ctx := context.Background()
	mockRepo := new(mocks.MockAccountRepository)
	mockRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockRepo.On("ClientExists", mock.Anything, mock.Anything).Return(nil)

	createAccount := NewCreateAccountHandler(mockRepo, &mocks.TestUoWManager{})

	response, err := createAccount.Handle(ctx, command)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	mockRepo.AssertExpectations(t)
	if len(mockRepo.Calls) != 2 {
		t.Errorf("expected 2 method calls, got %d", len(mockRepo.Calls))
	}

	t.Logf("\n"+
		"AccountID: %s",
		response.AccountID)
}
