package main

import (
	loadService "github.com/D1sordxr/simple-bank/outbox-processor/internal/application"
	loadStorage "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure"
	loadConfig "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app"
	loadLogger "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app/logger"
	loadSlogHandler "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app/logger/handlers"
	loadKafka "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/kafka"
	loadPostgres "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres"
	loadPostgresDAO "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres/dao"
	loadApp "github.com/D1sordxr/simple-bank/outbox-processor/internal/presentation"
)

// TODO: Dockerfile

func main() {
	cfg := loadConfig.NewConfig()

	slogLogger := loadSlogHandler.NewSlogLogger(cfg)
	logger := loadLogger.NewLogger(slogLogger)

	databaseConn := loadPostgres.NewConnection(&cfg.StorageConfig)

	outboxDAO := loadPostgresDAO.NewOutboxDAO(databaseConn)

	storage := loadStorage.NewStorage(
		outboxDAO, // write dao implementation
		outboxDAO, // read dao implementation
	)

	kafkaProducer := loadKafka.NewProducer(&cfg.KafkaConfig)

	processorService := loadService.NewOutboxProcessor(
		&cfg.AppConfig,
		logger,
		storage.OutboxCommandDAO,
		storage.OutboxQueryDAO,
		kafkaProducer,
	)

	app := loadApp.NewApp(processorService)
	app.RunApp()
}
