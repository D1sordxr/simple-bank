package transaction

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/interfaces"
	"github.com/segmentio/kafka-go"
)

type ProcessorConsumer struct {
	svc interfaces.ProcessTransaction
}

func NewTransactionProcessor(svc interfaces.ProcessTransaction) *ProcessorConsumer {
	return &ProcessorConsumer{svc: svc}
}

func (c *ProcessorConsumer) Process(ctx context.Context, msg kafka.Message) {
	data := dto.ProcessDTO{
		OutboxID: msg.Key,
		Data:     msg.Value,
	}

	err := c.svc.Handle(data)
	if err != nil {
		// log error
	}
}
