package account

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/account/vo"
	sharedExc "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/shared_exceptions"
	sharedVO "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/shared_vo"
	"time"
)

type Aggregate struct {
	AccountID sharedVO.UUID     // unique identifier for the account
	ClientID  sharedVO.UUID     // references client id
	Balance   vo.Balance        // current balance
	Currency  sharedVO.Currency // account currency (USD, EUR, RUB)
	Status    vo.Status         // status: active, closed, suspended
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(accountID sharedVO.UUID,
	clientID sharedVO.UUID,
	balance vo.Balance,
	currency sharedVO.Currency,
	status vo.Status,
) (Aggregate, error) {
	if accountID.IsNil() || clientID.IsNil() {
		return Aggregate{}, sharedExc.InvalidUUID
	}
	return Aggregate{
		AccountID: accountID,
		ClientID:  clientID,
		Balance:   balance,
		Currency:  currency,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
