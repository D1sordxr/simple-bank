package kafka

import (
	"context"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/kafka/config"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Writer *kafka.Writer
}

func NewProducer(config *config.KafkaConfig) *Producer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(config.Brokers...),
		Balancer: &kafka.LeastBytes{},
	}
	return &Producer{Writer: writer}
}

func (p *Producer) SendMessage(
	ctx context.Context,
	key,
	value []byte,
	topic string,
) error {
	return p.Writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Key:   key,
		Value: value,
	})
}

func (p *Producer) Close() error {
	return p.Writer.Close()
}
