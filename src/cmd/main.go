package main

import (
	loadApplicationServices "github.com/D1sordxr/simple-banking-system/internal/application"
	loadAccountService "github.com/D1sordxr/simple-banking-system/internal/application/account"
	loadAccountCommands "github.com/D1sordxr/simple-banking-system/internal/application/account/handlers"
	loadClientService "github.com/D1sordxr/simple-banking-system/internal/application/client"
	loadClientCommands "github.com/D1sordxr/simple-banking-system/internal/application/client/handlers"
	loadTransactionService "github.com/D1sordxr/simple-banking-system/internal/application/transaction"
	loadTransactionCommands "github.com/D1sordxr/simple-banking-system/internal/application/transaction/handlers"
	loadStorage "github.com/D1sordxr/simple-banking-system/internal/infrastructure"
	loadConfig "github.com/D1sordxr/simple-banking-system/internal/infrastructure/app"
	loadLogger "github.com/D1sordxr/simple-banking-system/internal/infrastructure/app/logger"
	loadPostgresConnection "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres"
	loadPostgresAccountRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/account"
	loadPostgresClientRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/client"
	loadPostgresEventRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/shared/event"
	loadPostgresOutboxRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/shared/outbox"
	loadPostgresTransactionRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/transaction"
	loadPostgresUoW "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/uow"
)

// TODO: Logging - finish createClientHandler logging and add logging for other services
// TODO: Errors - customize repository errors
// TODO: Dependencies - add for account and transaction use client as example
// TODO: Client - finish logic (event+outbox) -> infra (repo) -> presentation (grpc)
// TODO: Account - finish logic (event+outbox) -> infra (repo) -> presentation (grpc)
// TODO: Transaction - presentation (grpc)
// TODO: App - run app

// TODO: Outbox reader and Kafka producer service
// TODO: Kafka consumer service

func main() {
	cfg := loadConfig.NewConfig()

	logger := loadLogger.NewSlogLogger(cfg)

	databaseConn := loadPostgresConnection.NewConnection(&cfg.StorageConfig)

	uowManager := loadPostgresUoW.NewUoWManager(databaseConn)
	eventRepository := loadPostgresEventRepo.NewEventRepository(databaseConn)
	outboxRepository := loadPostgresOutboxRepo.NewOutboxRepository(databaseConn)

	clientRepository := loadPostgresClientRepo.NewClientRepository(databaseConn)
	accountRepository := loadPostgresAccountRepo.NewAccountRepository(databaseConn)
	transactionRepository := loadPostgresTransactionRepo.NewTransactionRepository(databaseConn)

	storage := loadStorage.NewStorage( // TODO: finish client and account repos
		uowManager,            // unitOfWork manager implementation
		eventRepository,       // event repository implementation
		outboxRepository,      // outbox repository implementation
		clientRepository,      // client repository implementation
		accountRepository,     // account repository implementation
		transactionRepository, // transaction repository implementation
	)

	clientDependencies := loadClientService.NewClientDependencies(
		logger,
		storage.UnitOfWork,
		storage.EventRepository,
		storage.OutboxRepository,
		storage.ClientRepository,
	)
	createClientCommand := loadClientCommands.NewCreateClientHandler(clientDependencies)
	updateClientCommand := loadClientCommands.NewUpdateClientHandler(clientDependencies) // TODO: updateClientHandler
	clientService := loadClientService.NewClientService(createClientCommand, updateClientCommand)

	createAccountCommand := loadAccountCommands.NewCreateAccountHandler(storage.AccountRepository, storage.UnitOfWork)
	getByIDAccountCommand := loadAccountCommands.NewGetByIDAccountHandler(storage.AccountRepository, storage.UnitOfWork)
	accountService := loadAccountService.NewAccountService(createAccountCommand, getByIDAccountCommand)

	createTransactionCommand := loadTransactionCommands.NewCreateTransactionHandler(storage.TransactionRepository, storage.UnitOfWork)
	transactionService := loadTransactionService.NewTransactionService(createTransactionCommand)

	applicationServices := loadApplicationServices.NewApplicationServices(
		clientService,      // client commands service
		accountService,     // account commands service
		transactionService, // transaction commands service
	)

	// TODO: gRPC := NewGRPCServer()

	// TODO: app := NewApp()
	// TODO: app.Run()
}
