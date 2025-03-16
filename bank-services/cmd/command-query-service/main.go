package main

import (
	pkgLog "github.com/D1sordxr/packages/log"
	loadApplicationServices "github.com/D1sordxr/simple-bank/bank-services/internal/application"
	loadAccountService "github.com/D1sordxr/simple-bank/bank-services/internal/application/account"
	loadAccountDeps "github.com/D1sordxr/simple-bank/bank-services/internal/application/account/dependencies"
	loadAccountHandlers "github.com/D1sordxr/simple-bank/bank-services/internal/application/account/handlers"
	loadClientService "github.com/D1sordxr/simple-bank/bank-services/internal/application/client"
	loadClientDeps "github.com/D1sordxr/simple-bank/bank-services/internal/application/client/dependencies"
	loadClientHandlers "github.com/D1sordxr/simple-bank/bank-services/internal/application/client/handlers"
	loadTransactionService "github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction"
	loadTransactionDeps "github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dependencies"
	loadTransactionHandlers "github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/handlers"
	loadAccountDomainSvc "github.com/D1sordxr/simple-bank/bank-services/internal/domain/account/services"
	loadTransactionDomainSvc "github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/services"
	loadStorage "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure"
	loadConfig "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app"
	loadLogger "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/logger"
	loadSlogLogger "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/logger/handlers"
	loadPostgresConnection "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres"
	loadPosgresExecutor "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/executor"
	loadPostgresAccountRepo "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/repositories/account"
	loadPostgresClientRepo "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/repositories/client"
	loadPostgresEventRepo "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/repositories/shared/event"
	loadPostgresOutboxRepo "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/repositories/shared/outbox"
	loadPostgresTransactionRepo "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/repositories/transaction"
	loadPostgresUoW "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/unit-of-work"
	loadApp "github.com/D1sordxr/simple-bank/bank-services/internal/presentation"
	loadGrpcServer "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc"
	loadGrpcServices "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/handlers"
	loadAccountGrpcService "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/handlers/account"
	loadClientGrpcService "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/handlers/client"
	loadTxGrpcService "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/handlers/transaction"
)

// TODO: UpdateCommands - implement client (executes event + outbox)

// TODO: Queries - add application logic for client and account, implement and use cache (projections) + DAOs
// TODO: Move queries to another app which depends on Redis

// TODO: Transaction (aggregate) - add reversal type support

// TODO: Redis for caching client and account data

// TODO: Money - rework float64 -> decimal.Decimal from shopspring library (optional)

func main() {
	cfg := loadConfig.NewConfig()

	slogLogger := loadSlogLogger.NewSlogLogger(cfg)
	logger := loadLogger.NewLogger(slogLogger)
	logV2 := pkgLog.Default()

	databasePool := loadPostgresConnection.NewPool(&cfg.StorageConfig)
	databaseExecutor := loadPosgresExecutor.NewExecutor(databasePool)

	unitOfWork := loadPostgresUoW.NewUnitOfWork(logger, databaseExecutor)

	eventRepository := loadPostgresEventRepo.NewEventRepository(databaseExecutor)
	outboxRepository := loadPostgresOutboxRepo.NewOutboxRepository(databaseExecutor)

	clientRepository := loadPostgresClientRepo.NewClientRepository(databaseExecutor)
	accountRepository := loadPostgresAccountRepo.NewAccountRepository(databaseExecutor)
	transactionRepository := loadPostgresTransactionRepo.NewTransactionRepository(databaseExecutor)

	storage := loadStorage.NewStorage(
		unitOfWork,            // unitOfWork implementation
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
	createClientCommand := loadClientHandlers.NewCreateClientHandler(clientDependencies)
	updateClientCommand := loadClientHandlers.NewUpdateClientHandler(clientDependencies) // TODO: updateClientHandler
	clientService := loadClientService.NewClientService(
		createClientCommand, // create client command implementation
		updateClientCommand, // update client command implementation
	)

	accountDependencies := loadAccountDeps.NewAccountDependencies(
		logger,
		storage.UnitOfWork,
		storage.EventRepository,
		storage.OutboxRepository,
		storage.AccountRepository,
	)
	createAccountCommand := loadAccountHandlers.NewCreateAccountHandler(accountDependencies)
	updateAccountCommand := loadAccountHandlers.NewUpdateAccountHandler(
		logV2,
		storage.UnitOfWork,
		storage.EventRepository,
		storage.OutboxRepository,
		new(loadAccountDomainSvc.UpdateDomainSvc),
	)
	accountService := loadAccountService.NewAccountService(
		createAccountCommand, // create account command implementation
		updateAccountCommand, // update account command implementation
	)

	transactionDependencies := loadTransactionDeps.NewTransactionDependencies(
		logger,
		storage.UnitOfWork,
		storage.EventRepository,
		storage.OutboxRepository,
		storage.TransactionRepository,
	)
	createTransactionCommand := loadTransactionHandlers.NewCreateTransactionHandler(transactionDependencies)
	updateTransactionCommand := loadTransactionHandlers.NewUpdateTransactionHandler(
		logV2,
		storage.UnitOfWork,
		storage.EventRepository,
		storage.OutboxRepository,
		new(loadTransactionDomainSvc.UpdateDomainSvc))
	transactionService := loadTransactionService.NewTransactionService(
		createTransactionCommand, // create transaction command implementation
		updateTransactionCommand, // update transaction command implementation
	)

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
