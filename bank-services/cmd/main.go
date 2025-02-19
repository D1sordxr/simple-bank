package main

import (
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
	loadStorage "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure"
	loadConfig "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app"
	loadLogger "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/logger"
	loadSlogLogger "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/logger/handlers"
	loadPostgresConnection "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres"
	loadPostgresAccountRepo "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/repositories/account"
	loadPostgresClientRepo "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/repositories/client"
	loadPostgresEventRepo "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/repositories/shared/event"
	loadPostgresOutboxRepo "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/repositories/shared/outbox"
	loadPostgresTransactionRepo "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/repositories/transaction"
	loadPostgresUoW "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/uow"
	loadApp "github.com/D1sordxr/simple-bank/bank-services/internal/presentation"
	loadGrpcServer "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc"
	loadGrpcServices "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/handlers"
	loadAccountGrpcService "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/handlers/account"
	loadClientGrpcService "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/handlers/client"
	loadTxGrpcService "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/handlers/transaction"
)

// TODO: UnitOfWork - change context values to "tx" and "batch" -> add repo support -> replace with storage.UoW
// TODO: Queries - add application logic for client and account, implement and use cache (projections) + DAOs
// TODO: Transaction - add reversal type support

// TODO: Workers...
// TODO: Outbox reader and Kafka producer service
// TODO: Kafka consumer services to process transaction
// TODO: Redis for caching client and account data

// TODO: Money - rework float64 -> decimal.Decimal from shopspring library (optional)
// TODO: Postgres - rework connection -> pool (optional)

func main() {
	cfg := loadConfig.NewConfig()

	slogLogger := loadSlogLogger.NewSlogLogger(cfg)
	logger := loadLogger.NewLogger(slogLogger)

	databasePool := loadPostgresConnection.NewPool(&cfg.StorageConfig)

	unitOfWork := loadPostgresUoW.NewUoW(databasePool)
	eventRepository := loadPostgresEventRepo.NewEventRepository(databasePool)
	outboxRepository := loadPostgresOutboxRepo.NewOutboxRepository(databasePool)

	clientRepository := loadPostgresClientRepo.NewClientRepository(databasePool)
	accountRepository := loadPostgresAccountRepo.NewAccountRepository(databasePool)
	transactionRepository := loadPostgresTransactionRepo.NewTransactionRepository(databasePool)

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
	getByIDAccountQuery := loadAccountHandlers.NewGetByIDAccountHandler(accountDependencies) // TODO: getByIDAccountQuery
	accountService := loadAccountService.NewAccountService(
		createAccountCommand, // create account command implementation
		getByIDAccountQuery,  // getByID account query implementation
	)

	transactionDependencies := loadTransactionDeps.NewTransactionDependencies(
		logger,
		storage.UnitOfWork,
		storage.EventRepository,
		storage.OutboxRepository,
		storage.TransactionRepository,
	)
	createTransactionCommand := loadTransactionHandlers.NewCreateTransactionHandler(transactionDependencies)
	transactionService := loadTransactionService.NewTransactionService(
		createTransactionCommand, // create transaction command implementation
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
