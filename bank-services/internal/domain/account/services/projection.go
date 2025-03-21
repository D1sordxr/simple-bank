package services

import (
	"encoding/json"
	"fmt"
	accountRoot "github.com/D1sordxr/simple-bank/bank-services/internal/domain/account"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/account/vo"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/consts"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/account"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/models"
	"github.com/google/uuid"
)

type AccountProjectionDomainSvc struct{}

func (AccountProjectionDomainSvc) ParseUpdateEvent(data []byte) (e account.UpdateEvent, err error) {
	err = json.Unmarshal(data, &e)
	if err != nil {
		return account.UpdateEvent{}, err
	}

	return e, nil
}

func (AccountProjectionDomainSvc) ConvertModelToProjection(model models.Account) accountRoot.Projection {
	return accountRoot.Projection{
		AccountID: model.ID.String(),
		Balance:   model.AvailableMoney,
		Status:    model.Status,
	}
}

func (AccountProjectionDomainSvc) UpdateProjection(oldProjection accountRoot.Projection, upd account.UpdateEvent) (accountRoot.Projection, error) {
	const op = "services.AccountProjectionDomainSvc.UpdateProjection"

	if oldProjection.AccountID != upd.AccountID {
		return accountRoot.Projection{}, fmt.Errorf("%s: %w", op, ErrDifferentAccountIDs)
	}

	switch oldProjection.Status {
	case vo.StatusClosed:
		return accountRoot.Projection{}, fmt.Errorf("%s: %w", op, ErrAccountHasStatusClosed)
	case vo.StatusSuspended:
		return accountRoot.Projection{}, fmt.Errorf("%s: %w", op, ErrAccountHasStatusSuspended)
	}

	switch upd.BalanceUpdateType {
	case consts.DebitBalanceUpdateType:
		if oldProjection.Balance < upd.Amount {
			return accountRoot.Projection{}, fmt.Errorf("%s: %w", op, ErrNotEnoughMoney)
		}
		oldProjection.Balance -= upd.Amount
	case consts.CreditBalanceUpdateType:
		oldProjection.Balance += upd.Amount
	}

	if oldProjection.Balance < 0 {
		return accountRoot.Projection{}, fmt.Errorf("%s: %w", op, ErrNegativeBalance)
	}

	return accountRoot.Projection{
		AccountID: oldProjection.AccountID,
		Balance:   oldProjection.Balance,
		Status:    upd.Status,
	}, nil
}

func (AccountProjectionDomainSvc) ConvertProjectionToModel(p accountRoot.Projection) (models.Account, error) {
	const op = "services.TransactionProjectionDomainSvc.ConvertProjectionToModel"

	id, err := uuid.Parse(p.AccountID)
	if err != nil {
		return models.Account{}, fmt.Errorf("%s: %w", op, FailedToParseUUID)
	}

	return models.Account{
		ID:             id,
		AvailableMoney: p.Balance,
		Status:         p.Status,
	}, nil
}
