package application

import (
	"context"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/commands"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/interfaces"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/interfaces/persistence"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/queries"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app"
	loadLogger "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app/logger"
	"os"
	"os/signal"
	"syscall"
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
	*loadLogger.Logger
	interfaces.OutboxCommand
	interfaces.OutboxQuery
	persistence.Producer
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
		Producer:        producer,
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
		logger.String("clientEmail", c.Email),
	)

	// TODO: uow.Begin()

	messages, err := p.OutboxQuery.FetchMessages(ctx, q)
	if err != nil {
		return err
	}

	for _, msg := range messages {
		err = p.Producer.SendMessage(ctx, []byte(msg.AggregateID), []byte(msg.MessagePayload))
		if err != nil {
			continue
		}
		c.ID = msg.OutboxID

		err = p.OutboxCommand.UpdateStatus(ctx, c)
		if err != nil {
			// TODO: Logger.Error()
		}
	}

	// TODO: uow.Commit()

	return nil
}

func (p *OutboxProcessor) processClientOutbox() {
	command := commands.OutboxCommand{
		Status: StatusProcessed,
	}
	query := queries.OutboxQuery{
		AggregateType: ClientAggregateType,
		Status:        StatusPending,
		Limit:         p.OutboxBatchSize,
	}

	ctx := context.Background()

	err := p.ProcessOutbox(ctx, command, query)
	if err != nil {
		// TODO: Logger.Error()
	}
}

func (p *OutboxProcessor) processAccountOutbox() {
	command := commands.OutboxCommand{
		Status: StatusProcessed,
	}
	query := queries.OutboxQuery{
		AggregateType: AccountAggregateType,
		Status:        StatusPending,
		Limit:         p.OutboxBatchSize,
	}

	ctx := context.Background()

	err := p.ProcessOutbox(ctx, command, query)
	if err != nil {
		// TODO: Logger.Error()
	}
}

func (p *OutboxProcessor) processTransactionOutbox() {
	command := commands.OutboxCommand{
		Status: StatusProcessed,
	}
	query := queries.OutboxQuery{
		AggregateType: TransactionAggregateType,
		Status:        StatusPending,
		Limit:         p.OutboxBatchSize,
	}

	ctx := context.Background()

	err := p.ProcessOutbox(ctx, command, query)
	if err != nil {
		// TODO: Logger.Error()
	}
}

func (p *OutboxProcessor) Run() {
	var err error
	errorsChannel := make(chan error, 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go p.processClientOutbox()
	go p.processAccountOutbox()
	go p.processTransactionOutbox()

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
		<-stop
		//s.Logger.Info("Stop signal received...")
		cancel()
	}()

	select {
	case <-ctx.Done():
		//s.Logger.Info("Stopping application...", s.Logger.String("reason", "stop signal"))
	case err = <-errorsChannel:
		//s.Logger.Error("Application encountered an error", s.Logger.String("error", err.Error()))
	}

	//s.Down()
	//s.Logger.Info("Gracefully stopped")
}
