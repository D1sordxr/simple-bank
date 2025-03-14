package interfaces

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/account"
)

type CompleteTransactionDomainSvc interface {
	UnmarshalData(dto dto.ProcessDTO) (account.UpdateEvents, error)
}
