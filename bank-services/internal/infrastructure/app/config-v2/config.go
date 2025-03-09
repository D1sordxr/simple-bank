package app

import (
	"github.com/D1sordxr/packages/kafka/consumer"
	"github.com/D1sordxr/packages/kafka/producer"
	"github.com/D1sordxr/packages/log"
	"github.com/D1sordxr/packages/postgres"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/kafka"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

const BasicConfigPath = "./configs/app/local.yaml"

type Config struct {
	App         App             `yaml:"app"`
	Logger      log.Config      `yaml:"logger"`
	Storage     postgres.Config `yaml:"storage"`
	Consumer    consumer.Config `yaml:"consumer"`
	Producer    producer.Config `yaml:"producer"`
	KafkaTopics kafka.Topics    `yaml:"kafka_topics"`
}

type App struct {
	Mode string `yaml:"mode"`
}

func NewConfig() *Config {
	var cfg Config

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = BasicConfigPath
	}

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}
