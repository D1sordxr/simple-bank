package interfaces

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/transaction"
	txRoot "github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/models"
)

type ProjectionDomainSvc interface {
	ParseUpdateEvent(data []byte) (e transaction.UpdateEvent, err error)
	ConvertModelToProjection(model models.TransactionModel) txRoot.Projection
	UpdateProjection(oldProjection txRoot.Projection, upd transaction.UpdateEvent) (txRoot.Projection, error)
	ConvertProjectionToModel(p txRoot.Projection) (models.TransactionModel, error)
}

type TransactionDAO interface {
	GetProjection(ctx context.Context, id string) (model models.TransactionModel, err error)
	UpdateProjection(ctx context.Context, model models.TransactionModel) error
}
