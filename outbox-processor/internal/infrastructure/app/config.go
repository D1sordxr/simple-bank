package app

import (
	kafkaConfig "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/kafka/config"
	storageConfig "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres/config"
)

type Config struct {
	AppConfig     App                         `yaml:"app"`
	StorageConfig storageConfig.StorageConfig `yaml:"storage"`
	KafkaConfig   kafkaConfig.KafkaConfig     `yaml:"kafka"`
}

type App struct {
	Mode            string `yaml:"mode"`
	OutboxBatchSize int    `yaml:"outbox_batch_size"`
}
