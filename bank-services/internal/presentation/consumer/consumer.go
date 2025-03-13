package consumer

import (
	"context"
	"github.com/D1sordxr/packages/kafka/consumer"
	"github.com/segmentio/kafka-go"
	"sync"
)

// This part of code could be omitted.
// Now consumer is used from "github.com/D1sordxr/packages/kafka/consumer"

type Processor interface {
	Process(ctx context.Context, msg kafka.Message) error
}

type Consumer struct {
	proc   Processor
	reader *kafka.Reader
}

func NewConsumer(config *consumer.Config, proc Processor) *Consumer {
	return &Consumer{
		// reader: consumer.NewConsumer(config),
		proc: proc,
	}
}

func (c *Consumer) ReadAndProcess(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			// log.Println("Consumer: получен сигнал завершения, останавливаем обработку...")
			return
		default:
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				// log.Printf("Ошибка при чтении сообщения: %v", err)
				continue
			}

			if err = c.proc.Process(ctx, msg); err != nil {
				// log.Printf("Ошибка при обработке сообщения: %v", err)
			}
		}
	}
}
