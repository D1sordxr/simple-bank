package services

import (
	"encoding/json"
	"fmt"
	eventTx "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/transaction"
	txRoot "github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/vo"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/models"
	"github.com/google/uuid"
)

type TransactionProjectionDomainSvc struct{}

func (TransactionProjectionDomainSvc) ParseUpdateEvent(data []byte) (e eventTx.UpdateEvent, err error) {
	err = json.Unmarshal(data, &e)
	if err != nil {
		return eventTx.UpdateEvent{}, err
	}

	return e, nil
}

func (TransactionProjectionDomainSvc) ConvertModelToProjection(model models.TransactionModel) txRoot.Projection {
	return txRoot.Projection{
		TransactionID: model.ID.String(),
		Status:        model.Status,
		FailureReason: model.FailureReason,
	}
}

func (TransactionProjectionDomainSvc) UpdateProjection(oldProjection txRoot.Projection, upd eventTx.UpdateEvent) (txRoot.Projection, error) {
	const op = "services.TransactionProjectionDomainSvc.UpdateProjection"

	if oldProjection.TransactionID != upd.TransactionID {
		return txRoot.Projection{}, fmt.Errorf("%s: %w", op, ErrDifferentTxIDs)
	}

	if oldProjection.Status == upd.Status {
		return txRoot.Projection{}, fmt.Errorf("%s: %w", op, ErrStatusesEqual)
	}

	var fReason *string
	if upd.FailureReason != "" {
		if oldProjection.Status == vo.StatusFailed {
			return txRoot.Projection{}, fmt.Errorf("%s: %w", op, ErrTxHasStatusFailed)
		}
		fReason = &upd.FailureReason
	}

	return txRoot.Projection{
		TransactionID: upd.TransactionID,
		Status:        upd.Status,
		FailureReason: fReason,
	}, nil
}

func (TransactionProjectionDomainSvc) ConvertProjectionToModel(p txRoot.Projection) (models.TransactionModel, error) {
	const op = "services.TransactionProjectionDomainSvc.ConvertProjectionToModel"

	id, err := uuid.Parse(p.TransactionID)
	if err != nil {
		return models.TransactionModel{}, fmt.Errorf("%s: %w", op, FailedToParseUUID)
	}

	return models.TransactionModel{
		ID:            id,
		Status:        p.Status,
		FailureReason: p.FailureReason,
	}, nil
}
