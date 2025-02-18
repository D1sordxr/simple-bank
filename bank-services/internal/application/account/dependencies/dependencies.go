package dependencies

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/shared/interfaces"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/account"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/outbox"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/logger"
)

type Dependencies struct {
	Logger            *logger.Logger
	UoWManager        interfaces.UnitOfWork
	EventRepository   event.Repository
	OutboxRepository  outbox.Repository
	AccountRepository account.Repository
}

func NewAccountDependencies(
	logger *logger.Logger,
	uow interfaces.UnitOfWork,
	event event.Repository,
	outbox outbox.Repository,
	repo account.Repository,
) *Dependencies {
	return &Dependencies{
		Logger:            logger,
		UoWManager:        uow,
		EventRepository:   event,
		OutboxRepository:  outbox,
		AccountRepository: repo,
	}
}
