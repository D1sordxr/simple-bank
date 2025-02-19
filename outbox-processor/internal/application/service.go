package application

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/commands"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/interfaces"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/interfaces/persistence"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/queries"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app"
	loadLogger "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	StatusPending   = "Pending"
	StatusProcessed = "Processed"
	StatusFailed    = "Failed"

	ClientAggregateType      = "Client"
	AccountAggregateType     = "Account"
	TransactionAggregateType = "Transaction"
)

type OutboxProcessor struct {
	Logger          *loadLogger.Logger
	OutboxCommand   interfaces.OutboxCommand
	OutboxQuery     interfaces.OutboxQuery
	KafkaProducer   persistence.Producer
	Ticker          time.Duration
	OutboxBatchSize int
}

func NewOutboxProcessor(
	cfg *app.App,
	log *loadLogger.Logger,
	c interfaces.OutboxCommand,
	q interfaces.OutboxQuery,
	producer persistence.Producer,
) *OutboxProcessor {
	return &OutboxProcessor{
		Logger:          log,
		OutboxCommand:   c,
		OutboxQuery:     q,
		KafkaProducer:   producer,
		Ticker:          cfg.Ticker,
		OutboxBatchSize: cfg.OutboxBatchSize,
	}
}

func (p *OutboxProcessor) ProcessOutbox(
	ctx context.Context,
	c commands.OutboxCommand,
	q queries.OutboxQuery,
) error {
	const op = "Service.OutboxProcessor.ProcessOutbox"

	logger := p.Logger
	log := logger.With(
		logger.String("operation", op),
		logger.Group("query",
			logger.String("aggregateType", q.AggregateType),
			logger.String("status", q.Status),
		),
	)

	log.Info("Starting processing outbox...")

	// TODO: uow-v0.Begin()

	messages, err := p.OutboxQuery.FetchMessages(ctx, q)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	for _, msg := range messages {
		err = p.KafkaProducer.SendMessage(ctx, []byte(msg.OutboxID), []byte(msg.MessagePayload))
		if err != nil {
			log.Error("Error producing kafka messages")
			return fmt.Errorf("%s, %w", op, err)
		}
		c.ID = msg.OutboxID

		if err = p.OutboxCommand.UpdateStatus(ctx, c); err != nil {
			c.Status = StatusFailed
			if err = p.OutboxCommand.UpdateStatus(ctx, c); err != nil {
				log.Error("Error updating outbox status", logger.String("outboxID", c.ID))
				return fmt.Errorf("%s, %w", op, err)
			}
			log.Error("Outbox received status failed", logger.String("outboxID", c.ID))
		}
	}

	// TODO: uow-v0.Commit()

	log.Info("Outbox processed successfully!")

	return nil
}

func (p *OutboxProcessor) processClientOutbox(ctx context.Context) error {
	command := commands.OutboxCommand{
		Status: StatusProcessed,
	}
	query := queries.OutboxQuery{
		AggregateType: ClientAggregateType,
		Status:        StatusPending,
		Limit:         p.OutboxBatchSize,
	}

	err := p.ProcessOutbox(ctx, command, query)
	if err != nil {
		return err
	}

	return nil
}

func (p *OutboxProcessor) processAccountOutbox(ctx context.Context) error {
	command := commands.OutboxCommand{
		Status: StatusProcessed,
	}
	query := queries.OutboxQuery{
		AggregateType: AccountAggregateType,
		Status:        StatusPending,
		Limit:         p.OutboxBatchSize,
	}

	err := p.ProcessOutbox(ctx, command, query)
	if err != nil {
		return err
	}

	return nil
}

func (p *OutboxProcessor) processTransactionOutbox(ctx context.Context) error {
	command := commands.OutboxCommand{
		Status: StatusProcessed,
	}
	query := queries.OutboxQuery{
		AggregateType: TransactionAggregateType,
		Status:        StatusPending,
		Limit:         p.OutboxBatchSize,
	}

	err := p.ProcessOutbox(ctx, command, query)
	if err != nil {
		return err
	}

	return nil
}

func (p *OutboxProcessor) Run() {
	wg := &sync.WaitGroup{}
	errorsChannel := make(chan error, 1)

	// Graceful shutdown context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Timer channel
	ticker := time.NewTicker(p.Ticker)
	defer ticker.Stop()

	// Signal handler for graceful shutdown
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
		<-stop
		p.Logger.Info("Stop signal received, shutting down...")
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			p.Logger.Info("Context cancelled, shutting down all processes...")
			wg.Wait()
			return
		case <-ticker.C:
			wg.Add(3)

			// Process Client Outbox
			go func() {
				defer wg.Done()

				ctxWithTimeout, timeoutCancel := context.WithTimeout(ctx, p.Ticker)
				defer timeoutCancel()

				if err := p.processClientOutbox(ctxWithTimeout); err != nil {
					errorsChannel <- err
				}
			}()

			// Process Account Outbox
			go func() {
				defer wg.Done()

				ctxWithTimeout, timeoutCancel := context.WithTimeout(ctx, p.Ticker)
				defer timeoutCancel()

				if err := p.processAccountOutbox(ctxWithTimeout); err != nil {
					errorsChannel <- err
				}
			}()

			// Process Transaction Outbox
			go func() {
				defer wg.Done()

				ctxWithTimeout, timeoutCancel := context.WithTimeout(ctx, p.Ticker)
				defer timeoutCancel()

				if err := p.processTransactionOutbox(ctxWithTimeout); err != nil {
					errorsChannel <- err
				}
			}()
		case err := <-errorsChannel:
			p.Logger.Error("Application encountered an error", p.Logger.String("error", err.Error()))

			// cancel()
		}
	}
}
