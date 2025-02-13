package account

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/account"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/models"
)

func ConvertAggregateToModel(account account.Aggregate) models.Account {
	return models.Account{
		ID:             account.AccountID.Value,
		ClientID:       account.ClientID.Value,
		AvailableMoney: account.Balance.AvailableMoney.Value,
		FrozenMoney:    account.Balance.FrozenMoney.Value,
		Currency:       account.Currency.Code,
		Status:         account.Status.CurrentStatus,
		CreatedAt:      account.CreatedAt,
		UpdatedAt:      account.UpdatedAt,
	}
}
