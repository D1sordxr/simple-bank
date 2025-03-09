package services

import (
	"encoding/json"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction"
)

type IProcessDomainSvc interface {
	ParseMessage(msg []byte) (transaction.Aggregate, error)
}

type ProcessDomainSvc struct{}

func (ProcessDomainSvc) ParseMessage(msg []byte) (transaction.Aggregate, error) {
	const op = "domain.transaction.ProcessDomainSvc.ParseMessage"

	var agg transaction.Aggregate

	err := json.Unmarshal(msg, &agg)
	if err != nil {
		return transaction.Aggregate{}, fmt.Errorf("%s: %w: %w", op, ErrParsingMsg, err)
	}

	return agg, nil
}
