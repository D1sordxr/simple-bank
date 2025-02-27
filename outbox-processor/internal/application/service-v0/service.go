package service_v0

import (
	"context"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/commands"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/interfaces"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/interfaces/persistence"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/queries"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app"
	loadLogger "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app/logger"
	"time"
)

type OutboxProcessor struct {
	Logger          *loadLogger.Logger
	OutboxCommand   interfaces.OutboxCommand
	OutboxQuery     interfaces.OutboxQuery
	UnitOfWork      persistence.UnitOfWork
	KafkaProducer   persistence.Producer
	Ticker          time.Duration
	OutboxBatchSize int
}

func NewOutboxProcessor(
	cfg *app.App,
	log *loadLogger.Logger,
	uow persistence.UnitOfWork,
	c interfaces.OutboxCommand,
	q interfaces.OutboxQuery,
	producer persistence.Producer,
) *OutboxProcessor {
	return &OutboxProcessor{
		Logger:          log,
		UnitOfWork:      uow,
		OutboxCommand:   c,
		OutboxQuery:     q,
		KafkaProducer:   producer,
		Ticker:          cfg.Ticker,
		OutboxBatchSize: cfg.OutboxBatchSize,
	}
}

func (p *OutboxProcessor) StartProcessingOutbox(ctx context.Context) error {
	const op = "Service.OutboxProcessor.StartProcessingOutbox"
	log := p.Logger.With(p.Logger.String("operation", op))

	ticker := time.NewTicker(p.Ticker)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("Shutting down outbox processor")
				return
			case <-ticker.C:
				// continue
			}

			// TODO: Query
			messages, err := p.OutboxQuery.FetchMessages(ctx, queries.OutboxQuery{})
			if err != nil {
				log.Error("Failed to fetch messages", "error", err.Error())
				continue
			}
			if len(messages) == 0 {
				log.Info("No messages to process")
				continue
			}

			ctx, err = p.UnitOfWork.BeginWithTxAndBatch(ctx)
			if err != nil {
				return
			}

			for _, msg := range messages {
				err = p.KafkaProducer.SendMessage(ctx, []byte(msg.OutboxID), []byte(msg.MessagePayload))
				if err != nil {
					log.Error("") // TODO: log
					continue
				}

				err = p.OutboxCommand.UpdateStatus(ctx, commands.OutboxCommand{})
				if err != nil {
					log.Error("") // TODO: log
					continue
				}
			}
		}
	}()

	return nil
}
