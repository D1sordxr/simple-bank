package consumer

import (
	"context"
	sharedInterfaces "github.com/D1sordxr/simple-bank/bank-services/internal/application/shared/interfaces"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/segmentio/kafka-go"
)

type Handler struct {
	svc sharedInterfaces.MessageProcessor
}

func NewHandler(svc sharedInterfaces.MessageProcessor) *Handler {
	return &Handler{svc: svc}
}

func (c *Handler) Handle(ctx context.Context, msg kafka.Message) error {
	data := dto.ProcessDTO{
		Key:   msg.Key,
		Value: msg.Value,
	}

	err := c.svc.Process(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
