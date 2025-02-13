package infrastructure

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/outbox"
)

type Storage struct {
	OutboxRepository *outbox.Repository
}

func NewStorage(outboxRepo *outbox.Repository) *Storage {
	return &Storage{OutboxRepository: outboxRepo}
}
