package main

import (
	loadApplicationServices "github.com/D1sordxr/simple-banking-system/internal/application"
	loadAccountService "github.com/D1sordxr/simple-banking-system/internal/application/account"
	loadAccountDeps "github.com/D1sordxr/simple-banking-system/internal/application/account/commands"
	loadAccountCommands "github.com/D1sordxr/simple-banking-system/internal/application/account/handlers"
	loadClientService "github.com/D1sordxr/simple-banking-system/internal/application/client"
	loadClientDeps "github.com/D1sordxr/simple-banking-system/internal/application/client/commands"
	loadClientCommands "github.com/D1sordxr/simple-banking-system/internal/application/client/handlers"
	loadTransactionService "github.com/D1sordxr/simple-banking-system/internal/application/transaction"
	loadTransactionDeps "github.com/D1sordxr/simple-banking-system/internal/application/transaction/commands"
	loadTransactionCommands "github.com/D1sordxr/simple-banking-system/internal/application/transaction/handlers"
	loadStorage "github.com/D1sordxr/simple-banking-system/internal/infrastructure"
	loadConfig "github.com/D1sordxr/simple-banking-system/internal/infrastructure/app"
	loadLogger "github.com/D1sordxr/simple-banking-system/internal/infrastructure/app/logger"
	loadSlogLogger "github.com/D1sordxr/simple-banking-system/internal/infrastructure/app/logger/handlers"
	loadPostgresConnection "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres"
	loadPostgresAccountRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/account"
	loadPostgresClientRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/client"
	loadPostgresEventRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/shared/event"
	loadPostgresOutboxRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/shared/outbox"
	loadPostgresTransactionRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/transaction"
	loadPostgresUoW "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/uow"
	loadApp "github.com/D1sordxr/simple-banking-system/internal/presentation"
	loadGrpcServer "github.com/D1sordxr/simple-banking-system/internal/presentation/grpc"
	loadGrpcServices "github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/handlers"
	loadAccountGrpcService "github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/handlers/account"
	loadClientGrpcService "github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/handlers/client"
	loadTxGrpcService "github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/handlers/transaction"
)

// TODO: Transaction - add reversal type support

// TODO: Redis for caching client and account data
// TODO: Outbox reader and Kafka producer service
// TODO: Kafka consumer services to process transaction and caching client and account data

func main() {
	cfg := loadConfig.NewConfig()

	slogLogger := loadSlogLogger.NewSlogLogger(cfg)
	logger := loadLogger.NewLogger(slogLogger)

	databaseConn := loadPostgresConnection.NewConnection(&cfg.StorageConfig)

	uowManager := loadPostgresUoW.NewUoWManager(databaseConn)
	eventRepository := loadPostgresEventRepo.NewEventRepository(databaseConn)
	outboxRepository := loadPostgresOutboxRepo.NewOutboxRepository(databaseConn)

	clientRepository := loadPostgresClientRepo.NewClientRepository(databaseConn)
	accountRepository := loadPostgresAccountRepo.NewAccountRepository(databaseConn)
	transactionRepository := loadPostgresTransactionRepo.NewTransactionRepository(databaseConn)

	storage := loadStorage.NewStorage(
		uowManager,            // unitOfWork manager implementation
		eventRepository,       // event repository implementation
		outboxRepository,      // outbox repository implementation
		clientRepository,      // client repository implementation
		accountRepository,     // account repository implementation
		transactionRepository, // transaction repository implementation
	)

	clientDependencies := loadClientDeps.NewClientDependencies(
		logger,
		storage.UnitOfWork,
		storage.EventRepository,
		storage.OutboxRepository,
		storage.ClientRepository,
	)
	createClientCommand := loadClientCommands.NewCreateClientHandler(clientDependencies)
	updateClientCommand := loadClientCommands.NewUpdateClientHandler(clientDependencies) // TODO: updateClientHandler
	clientService := loadClientService.NewClientService(createClientCommand, updateClientCommand)

	accountDependencies := loadAccountDeps.NewAccountDependencies(
		logger,
		storage.UnitOfWork,
		storage.EventRepository,
		storage.OutboxRepository,
		storage.AccountRepository,
	)
	createAccountCommand := loadAccountCommands.NewCreateAccountHandler(accountDependencies)
	getByIDAccountCommand := loadAccountCommands.NewGetByIDAccountHandler(accountDependencies) // TODO: rework getByIDAccountCommand
	accountService := loadAccountService.NewAccountService(createAccountCommand, getByIDAccountCommand)

	transactionDependencies := loadTransactionDeps.NewTransactionDependencies(
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

	grpcClientService := loadClientGrpcService.NewClientGrpcService(applicationServices.ClientService)
	grpcAccountService := loadAccountGrpcService.NewAccountGrpcService(applicationServices.AccountService)
	grpcTransactionService := loadTxGrpcService.NewTransactionGrpcService(applicationServices.TransactionService)

	grpcServices := loadGrpcServices.NewGrpcServices(
		grpcClientService,      // client grpc service implementation
		grpcAccountService,     // account grpc service implementation
		grpcTransactionService, // transaction grpc service implementation
	)

	grpcServer := loadGrpcServer.NewGrpcServer(&cfg.GrpcConfig, logger, grpcServices)

	app := loadApp.NewApp(grpcServer)

	app.RunApp()
}
