package interfaces

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/commands"
	eventPkg "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/account"
)

type UpdateDomainSvc interface {
	CreateUpdateEvent(upd account.UpdateEvent) (eventPkg.Event, error)
	ConvertCommandToUpdEvent(c commands.UpdateAccountCommand) account.UpdateEvent
}
