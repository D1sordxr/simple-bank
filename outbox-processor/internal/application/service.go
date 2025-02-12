package application

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/persistence"
	"github.com/D1sordxr/simple-banking-system/internal/domain/outbox"
)

type OutboxProcessor struct {
	outbox.Repository
	persistence.UoWManager
	persistence.Producer
}

func NewOutboxProcessor(repo outbox.Repository, producer persistence.Producer) *OutboxProcessor {
	return &OutboxProcessor{Repository: repo, Producer: producer}
}

func (p *OutboxProcessor) ProcessOutbox(ctx context.Context) error {
	uow := p.UoWManager.GetUoW()
	tx, err := uow.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = uow.Rollback()
			// TODO: add logging for panic
			panic(r)
		}
		if err != nil {
			_ = uow.Rollback()
		}
	}()

	messages, err := p.Repository.FetchPendingMessages(ctx, tx, outbox.BatchSize)
	if err != nil {
		return err
	}

	for _, msg := range messages {
		err = p.Producer.SendMessage(ctx, []byte(msg.AggregateID), []byte(msg.MessagePayload))
		if err != nil {
			continue
		}

		err = p.Repository.MarkAsProcessed(ctx, tx, msg.OutboxID)
		if err != nil {
			// TODO: add logging for failed messages
		}
	}

	if err = uow.Commit(); err != nil {
		return err
	}
	return nil
}
