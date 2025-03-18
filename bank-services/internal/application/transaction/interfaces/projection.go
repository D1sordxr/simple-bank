package interfaces

import "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/transaction"

type ProjectionDomainSvc interface {
	ParseUpdateEvent(data []byte) (e transaction.UpdateEvent, err error)
}
