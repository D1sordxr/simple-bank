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

// TODO: Logging - add logging for services
// TODO: Errors - customize repository errors
// TODO: Dependencies - add for account and transaction use client as example
// TODO: Client - presentation (grpc)
// TODO: Account - finish logic (event+outbox) -> infra (repo) -> presentation (grpc)
// TODO: Transaction - presentation (grpc)
// TODO: App - run app

// TODO: Redis for caching client and account data
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

	storage := loadStorage.NewStorage( // TODO: finish account repo
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

	accountDependencies := loadAccountService.NewAccountDependencies(
		logger,
		storage.UnitOfWork,
		storage.EventRepository,
		storage.OutboxRepository,
		storage.AccountRepository,
	)
	createAccountCommand := loadAccountCommands.NewCreateAccountHandler(accountDependencies)
	getByIDAccountCommand := loadAccountCommands.NewGetByIDAccountHandler(accountDependencies) // TODO: rework getByIDAccountCommand
	accountService := loadAccountService.NewAccountService(createAccountCommand, getByIDAccountCommand)

	transactionDependencies := loadTransactionService.NewTransactionDependencies(
		logger,
		storage.UnitOfWork,
		storage.EventRepository,
		storage.OutboxRepository,
		storage.TransactionRepository,
	)
	createTransactionCommand := loadTransactionCommands.NewCreateTransactionHandler(transactionDependencies)
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
