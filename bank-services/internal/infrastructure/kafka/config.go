package kafka

import (
	"github.com/D1sordxr/packages/kafka/consumer"
	"github.com/D1sordxr/packages/kafka/producer"
)

type Config struct {
	Consumer       consumer.Config `yaml:"consumer"`
	ConsumerTopics ConsumerTopics  `yaml:"consumer_topics"`
	Producer       producer.Config `yaml:"producer"`
	ProducerTopics ProducerTopics  `yaml:"producer_topics"`
}
