package producer

import (
	"github.com/D1sordxr/packages/kafka/producer"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/kafka"
)

type Config struct {
	PkgKafka producer.Config `yaml:"pkg_kafka"`
	Topics   kafka.Topics    `yaml:"topics"`
}
