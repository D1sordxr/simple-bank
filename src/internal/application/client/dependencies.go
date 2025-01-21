package client

import (
	"github.com/D1sordxr/simple-banking-system/internal/application/persistence"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
)

type Dependencies struct {
	UoWManager       persistence.UoWManager
	EventRepository  event.Repository
	OutboxRepository outbox.Repository
	ClientRepository client.Repository
}

func NewClientDependencies(
	uow persistence.UoWManager,
	event event.Repository,
	outbox outbox.Repository,
	repo client.Repository,
) *Dependencies {
	return &Dependencies{
		UoWManager:       uow,
		EventRepository:  event,
		OutboxRepository: outbox,
		ClientRepository: repo,
	}
}
