package dependencies

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/shared/interfaces"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/client"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/outbox"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/logger"
)

type Dependencies struct {
	Logger           *logger.Logger
	UoWManager       interfaces.UoWManager
	EventRepository  event.Repository
	OutboxRepository outbox.Repository
	ClientRepository client.Repository
}

func NewClientDependencies(
	logger *logger.Logger,
	uow interfaces.UoWManager,
	event event.Repository,
	outbox outbox.Repository,
	repo client.Repository,
) *Dependencies {
	return &Dependencies{
		Logger:           logger,
		UoWManager:       uow,
		EventRepository:  event,
		OutboxRepository: outbox,
		ClientRepository: repo,
	}
}
