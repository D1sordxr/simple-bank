package consumer

import (
	"context"
	"github.com/D1sordxr/packages/kafka/consumer"
	"github.com/segmentio/kafka-go"
	"log"
)

type Consumer struct {
	*kafka.Reader
}

func NewConsumer(config *consumer.Config) *Consumer {
	return &Consumer{
		Reader: consumer.NewConsumer(config),
	}
}

func (c *Consumer) Run() {
	for {
		msg, err := c.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Ошибка при чтении сообщения: %v", err)
			continue
		}

		// Направляем сообщение в обработчик
		go handleMessage(msg.Value)
	}
}
