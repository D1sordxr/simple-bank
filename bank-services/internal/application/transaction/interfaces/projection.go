package interfaces

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/transaction"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/models"
)

type ProjectionDomainSvc interface {
	ParseUpdateEvent(data []byte) (e transaction.UpdateEvent, err error)
}

type TransactionDAO interface {
	GetTransaction(ctx context.Context, id string) (model models.TransactionModel, err error)
}
