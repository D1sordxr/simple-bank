package interfaces

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/commands"
	eventPkg "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/transaction"
)

type UpdateDomainSvc interface {
	CreateUpdateEvent(upd transaction.UpdateEvent) (eventPkg.Event, error)
	ConvertCommandToUpdEvent(c commands.UpdateTransactionCommand) transaction.UpdateEvent
}
