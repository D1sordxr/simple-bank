package dependencies

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/shared/interfaces"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/outbox"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/logger"
)

type Dependencies struct {
	Logger                *logger.Logger
	UoWManager            interfaces.UoWManager
	EventRepository       event.Repository
	OutboxRepository      outbox.Repository
	TransactionRepository transaction.Repository
}

func NewTransactionDependencies(
	logger *logger.Logger,
	uow interfaces.UoWManager,
	event event.Repository,
	outbox outbox.Repository,
	repo transaction.Repository,
) *Dependencies {
	return &Dependencies{
		Logger:                logger,
		UoWManager:            uow,
		EventRepository:       event,
		OutboxRepository:      outbox,
		TransactionRepository: repo,
	}
}
