package services

import (
	"encoding/json"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction"
)

type ProcessDomainSvc struct{}

func (ProcessDomainSvc) ParseMessage(dto dto.ProcessDTO) (string, transaction.Aggregate, error) {
	const op = "domain.transaction.ProcessDomainSvc.ParseMessage"

	var agg transaction.Aggregate

	err := json.Unmarshal(dto.Value, &agg)
	if err != nil {
		return "", transaction.Aggregate{}, fmt.Errorf("%s: %w: %w", op, ErrParsingMsg, err)
	}

	outboxID := string(dto.Key)

	return outboxID, agg, nil
}

func (ProcessDomainSvc) MarshalMessage(msg any) ([]byte, error) {
	return json.Marshal(msg)
}
