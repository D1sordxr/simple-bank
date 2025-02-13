package main

import (
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app"
	logger2 "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app/logger"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app/logger/handlers"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/kafka"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres/dao"
)

func main() {
	cfg := app.NewConfig()

	slogLogger := handlers.NewSlogLogger(cfg)
	logger := logger2.NewLogger(slogLogger)
	logger.Error("app logging in not implemented")

	databaseConn := postgres.NewConnection(&cfg.StorageConfig)

	outboxDAO := dao.NewOutboxDAO(databaseConn)

	storage := infrastructure.NewStorage(
		outboxDAO,
		outboxDAO,
	)

	kafkaProducer := kafka.NewProducer(&cfg.KafkaConfig)

	processorService := application.NewOutboxProcessor(
		storage.OutboxCommandDAO,
		storage.OutboxQueryDAO,
		kafkaProducer,
	)
}
