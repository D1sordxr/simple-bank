package interfaces

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/commands"
	eventPkg "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
)

type UpdateDomainSvc interface {
	CreateUpdateEvent(c commands.UpdateTransactionCommand) (eventPkg.Event, error)
}
