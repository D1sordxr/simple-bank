package infrastructure

import (
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/interfaces"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/interfaces/persistence"
)

type Storage struct {
	UnitOfWork       persistence.UnitOfWork
	OutboxCommandDAO interfaces.OutboxCommand
	OutboxQueryDAO   interfaces.OutboxQuery
}

func NewStorage(
	uow persistence.UnitOfWork,
	c interfaces.OutboxCommand,
	q interfaces.OutboxQuery,
) *Storage {
	return &Storage{
		UnitOfWork:       uow,
		OutboxCommandDAO: c,
		OutboxQueryDAO:   q,
	}
}
