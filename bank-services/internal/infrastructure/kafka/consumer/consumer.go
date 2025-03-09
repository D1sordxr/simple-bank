package consumer

import (
	"github.com/D1sordxr/packages/kafka/consumer"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/kafka"
)

type Config struct {
	PkgKafka consumer.Config `yaml:"pkg_kafka"`
	Topics   kafka.Topics    `yaml:"topics"`
}
