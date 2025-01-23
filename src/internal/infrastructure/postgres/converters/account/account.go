package account

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/account"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/models"
)

func ConvertAggregateToModel(account account.Aggregate) models.Account {
	return models.Account{
		ID:        account.AccountID.Value,
		ClientID:  account.ClientID.Value,
		Balance:   account.Balance.AvailableMoney.Value,
		Currency:  account.Currency.Code,
		Status:    account.Status.CurrentStatus,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
