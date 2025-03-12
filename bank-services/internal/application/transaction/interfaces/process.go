package interfaces

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction"
)

type ProcessTransaction interface {
	Handle(ctx context.Context, dto dto.ProcessDTO) error
}

type ProcessTransactionDAO interface {
	SetProcessed(ctx context.Context, msgID string) error
	IsProcessed(ctx context.Context, msgID string) (bool, error)
}

type ProcessDomainSvc interface {
	ParseMessage(dto dto.ProcessDTO) (string, transaction.Aggregate, error)
	MarshalMessage(msg any) ([]byte, error)
}
