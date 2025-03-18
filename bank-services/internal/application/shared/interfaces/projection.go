package interfaces

import "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"

type ProjectionDomainSvc interface {
	ParseEvent(data []byte) (e event.Event, err error)
}
