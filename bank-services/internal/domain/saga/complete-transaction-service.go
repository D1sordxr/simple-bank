package saga

import (
	"encoding/json"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/account"
)

type CompleteTransactionDomainSvc struct{}

func (CompleteTransactionDomainSvc) UnmarshalData(dto dto.ProcessDTO) (account.UpdateEvents, error) {
	const op = "domain.saga.CompleteTransactionDomainSvc.UnmarshalData"

	var data account.UpdateEvents
	err := json.Unmarshal(dto.Value, &data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return data, nil
}
