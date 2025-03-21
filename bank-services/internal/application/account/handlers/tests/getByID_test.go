package tests

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/commands"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/dependencies"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/handlers"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/account"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/account/vo"
	sharedVO "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/shared_vo"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/mocks"
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
			AvailableMoney: sharedVO.Money{Value: 42.12},
			FrozenMoney:    sharedVO.Money{Value: 12.42},
		},
		Currency:  sharedVO.Currency{Code: "USD"},
		Status:    vo.Status{CurrentStatus: "active"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	deps := new(dependencies.Dependencies)
	getByIDAccount := handlers.NewGetByIDAccountHandler(deps)

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
