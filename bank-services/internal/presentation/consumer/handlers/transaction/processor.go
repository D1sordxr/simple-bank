package transaction

import (
	"context"
	sharedInterfaces "github.com/D1sordxr/simple-bank/bank-services/internal/application/shared/interfaces"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/segmentio/kafka-go"
)

type MessageProcessor struct {
	svc sharedInterfaces.MessageProcessor
}

func NewTransactionProcessor(svc sharedInterfaces.MessageProcessor) *MessageProcessor {
	return &MessageProcessor{svc: svc}
}

func (c *MessageProcessor) Process(ctx context.Context, msg kafka.Message) error {
	data := dto.ProcessDTO{
		OutboxID: msg.Key,
		Data:     msg.Value,
	}

	err := c.svc.Handle(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
