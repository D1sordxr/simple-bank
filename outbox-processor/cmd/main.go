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
	loadExecutor "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres/executor"
	loadUoW "github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres/unit-of-work"
	loadApp "github.com/D1sordxr/simple-bank/outbox-processor/internal/presentation"
)

func main() {
	cfg := loadConfig.NewConfig()

	slogLogger := loadSlogHandler.NewSlogLogger(cfg)
	logger := loadLogger.NewLogger(slogLogger)

	_ = loadPostgres.NewConnection(&cfg.StorageConfig)
	databasePool := loadPostgres.NewPool(&cfg.StorageConfig)
	databaseExecutor := loadExecutor.NewExecutor(databasePool)

	unitOfWork := loadUoW.NewUnitOfWork(logger, databaseExecutor)

	outboxDAO := loadPostgresDAO.NewOutboxDAO(databaseExecutor)

	storage := loadStorage.NewStorage(
		unitOfWork, // unitOfWork implementation
		outboxDAO,  // write dao implementation
		outboxDAO,  // read dao implementation
	)

	kafkaProducer := loadKafka.NewProducer(&cfg.KafkaConfig)

	processorService := loadService.NewOutboxProcessor(
		&cfg.AppConfig,
		logger,
		storage.UnitOfWork,
		storage.OutboxCommandDAO,
		storage.OutboxQueryDAO,
		kafkaProducer,
	)

	app := loadApp.NewApp(processorService)
	app.RunApp()
}
