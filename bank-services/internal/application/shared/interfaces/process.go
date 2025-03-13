package interfaces

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
)

type MessageProcessor interface {
	Process(ctx context.Context, dto dto.ProcessDTO) error
}
