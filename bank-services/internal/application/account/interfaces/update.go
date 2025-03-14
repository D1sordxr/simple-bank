package interfaces

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/commands"
	eventPkg "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
)

type UpdateDomainSvc interface {
	CreateUpdateEvent(c commands.UpdateAccountCommand) (eventPkg.Event, error)
}
