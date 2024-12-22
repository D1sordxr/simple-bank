package account

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/account/vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_exceptions"
	vo2 "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_vo"
	"github.com/google/uuid"
	"time"
)

type Aggregate struct {
	AccountID uuid.UUID    // unique identifier for the account
	ClientID  uuid.UUID    // references client
	Balance   vo.Balance   // current balance
	Currency  vo2.Currency // account currency (USD, EUR, RUB)
	Status    vo.Status    // status: active, closed, suspended
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(accountID uuid.UUID,
	clientID uuid.UUID,
	balance vo.Balance,
	currency vo2.Currency) (Aggregate, error) {
	if accountID == uuid.Nil || clientID == uuid.Nil {
		return Aggregate{}, shared_exceptions.InvalidUUID
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
