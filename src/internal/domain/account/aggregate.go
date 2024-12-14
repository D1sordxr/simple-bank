package account

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/account/vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/exceptions"
	"github.com/google/uuid"
	"time"
)

type Aggregate struct {
	AccountID uuid.UUID   // unique identifier for the account
	ClientID  uuid.UUID   // references client
	Balance   vo.Balance  // current balance
	Currency  vo.Currency // account currency (USD, EUR, RUB)
	Status    vo.Status   // status: active, closed, suspended
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(accountID uuid.UUID,
	clientID uuid.UUID,
	balance vo.Balance,
	currency vo.Currency) (Aggregate, error) {
	if accountID == uuid.Nil || clientID == uuid.Nil {
		return Aggregate{}, exceptions.InvalidUUID
	}
	return Aggregate{
		AccountID: accountID,
		ClientID:  clientID,
		Balance:   balance,
		Currency:  currency,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
