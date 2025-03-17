package saga

import (
	"encoding/json"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/account"
)

type CompleteTransactionDomainSvc struct{}

func (CompleteTransactionDomainSvc) UnmarshalData(
	dto dto.ProcessDTO,
) (txID string, updEvents account.UpdateEvents, err error) {
	const op = "domain.saga.CompleteTransactionDomainSvc.UnmarshalData"

	err = json.Unmarshal(dto.Value, &updEvents)
	if err != nil {
		return "", nil, fmt.Errorf("%s: %w", op, err)
	}

	txID = updEvents[0].TransactionID

	return txID, updEvents, nil
}
