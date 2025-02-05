package infrastructure

import (
	"github.com/D1sordxr/simple-banking-system/internal/application/persistence"
	"github.com/D1sordxr/simple-banking-system/internal/domain/account"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction"
)

type Storage struct {
	UnitOfWork            persistence.UoWManager
	EventRepository       event.Repository
	OutboxRepository      outbox.Repository
	ClientRepository      client.Repository
	AccountRepository     account.Repository
	TransactionRepository transaction.Repository
}

func NewStorage(uow persistence.UoWManager,
	event event.Repository,
	outbox outbox.Repository,
	client client.Repository,
	account account.Repository,
	transaction transaction.Repository) *Storage {
	return &Storage{
		UnitOfWork:            uow,
		EventRepository:       event,
		OutboxRepository:      outbox,
		ClientRepository:      client,
		AccountRepository:     account,
		TransactionRepository: transaction,
	}
}
