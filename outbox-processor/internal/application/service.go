package application

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/commands"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/interfaces"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/interfaces/persistence"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/queries"
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
	interfaces.OutboxCommand
	interfaces.OutboxQuery
	persistence.UoWManager
	persistence.Producer
}

func NewOutboxProcessor(
	c interfaces.OutboxCommand,
	q interfaces.OutboxQuery,
	uow persistence.UoWManager,
	producer persistence.Producer,
) *OutboxProcessor {
	return &OutboxProcessor{
		OutboxCommand: c,
		OutboxQuery:   q,
		UoWManager:    uow,
		Producer:      producer,
	}
}

func (p *OutboxProcessor) ProcessOutbox(
	ctx context.Context,
	c commands.OutboxCommand,
	q queries.OutboxQuery,
) error {

	// TODO: Logger.Info()

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

		err = p.OutboxCommand.UpdateStatus(ctx, c)
		if err != nil {
			// TODO: Logger.Error()
		}
	}

	// TODO: uow.Commit()

	return nil
}
