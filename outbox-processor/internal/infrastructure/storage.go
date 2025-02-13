package infrastructure

import (
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/interfaces"
)

type Storage struct {
	OutboxCommandDAO interfaces.OutboxCommand
	OutboxQueryDAO   interfaces.OutboxQuery
}

func NewStorage(c interfaces.OutboxCommand, q interfaces.OutboxQuery) *Storage {
	return &Storage{
		OutboxCommandDAO: c,
		OutboxQueryDAO:   q,
	}
}
