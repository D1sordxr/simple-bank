package services

import (
	"encoding/json"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/transaction"
)

type TransactionProjectionDomainSvc struct{}

func (TransactionProjectionDomainSvc) ParseUpdateEvent(data []byte) (e transaction.UpdateEvent, err error) {
	err = json.Unmarshal(data, &e)
	if err != nil {
		return transaction.UpdateEvent{}, err
	}

	return e, nil
}
