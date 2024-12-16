package handlers

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/account/commands"
	"github.com/D1sordxr/simple-banking-system/internal/domain/account"
	"github.com/D1sordxr/simple-banking-system/internal/domain/account/vo"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/mocks"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestSuccessGetByIDAccountHandler(t *testing.T) {
	const fixedID = "077c39e1-d2f3-435a-9b47-49eac55bc1d3"

	command := commands.GetByIDAccountCommand{
		AccountID: fixedID,
	}

	ctx := context.Background()
	mockRepo := new(mocks.MockAccountRepository)
	mockRepo.On("GetByID", mock.Anything, mock.Anything).Return(account.Aggregate{
		Balance: vo.Balance{
			AvailableMoney: vo.Money{Value: 42.12},
			FrozenMoney:    vo.Money{Value: 12.42},
		},
		Currency:  vo.Currency{Code: "USD"},
		Status:    vo.Status{CurrentStatus: "active"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	getByIDAccount := NewGetByIDAccountHandler(mockRepo, &mocks.TestUoWManager{})

	response, err := getByIDAccount.Handle(ctx, command)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	mockRepo.AssertExpectations(t)
	if len(mockRepo.Calls) != 1 {
		t.Errorf("expected 1 method calls, got %d", len(mockRepo.Calls))
	}

	t.Logf("\n"+
		"AvailableMoney: %v\n"+
		"FrozenMoney: %v\n"+
		"Currency: %s\n"+
		"Status: %s\n"+
		"CreatedAt: %s\n"+
		"UpdatedAt: %s\n",
		response.AvailableMoney,
		response.FrozenMoney,
		response.Currency,
		response.Status,
		response.CreatedAt,
		response.UpdatedAt,
	)
}
