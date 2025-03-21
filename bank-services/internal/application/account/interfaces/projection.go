package interfaces

import (
	"context"
	accountRoot "github.com/D1sordxr/simple-bank/bank-services/internal/domain/account"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/account"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/models"
)

type ProjectionDomainSvc interface {
	ParseUpdateEvent(data []byte) (e account.UpdateEvent, err error)
	ConvertModelToProjection(model models.Account) accountRoot.Projection
	UpdateProjection(oldProjection accountRoot.Projection, upd account.UpdateEvent) (accountRoot.Projection, error)
	ConvertProjectionToModel(p accountRoot.Projection) (models.Account, error)
}

type AccountDAO interface {
	GetProjection(ctx context.Context, id string) (model models.Account, err error)
	UpdateProjection(ctx context.Context, model models.Account) error
}
