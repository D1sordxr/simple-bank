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
	loadPostgresConnection "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres"
	loadPostgresAccountRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/account"
	loadPostgresClientRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/client"
	loadPostgresTransactionRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/transaction"
	loadPostgresUoW "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/uow"
)

// TODO: EventStore - shared repo with SaveEvent() method
// TODO: Outbox - shared repo with SaveOutboxEvent() method
// TODO: Client - finish logic (event+outbox) -> infra (repo) -> presentation (grpc)
// TODO: Account - finish logic (event+outbox) -> infra (repo) -> presentation (grpc)
// TODO: Transaction - presentation (grpc)
// TODO: App - run app

// TODO: Outbox reader and Kafka producer service
// TODO: Kafka consumer service

func main() {
	cfg := loadConfig.NewConfig()

	databaseConn := loadPostgresConnection.NewConnection(&cfg.StorageConfig)

	uowManager := loadPostgresUoW.NewUoWManager(databaseConn)
	clientRepository := loadPostgresClientRepo.NewClientRepository(databaseConn)
	accountRepository := loadPostgresAccountRepo.NewAccountRepository(databaseConn)
	transactionRepository := loadPostgresTransactionRepo.NewTransactionRepository(databaseConn)

	storage := loadStorage.NewStorage( // TODO: finish client and account repos
		uowManager,            // uow implementation
		clientRepository,      // client repository implementation
		accountRepository,     // account repository implementation
		transactionRepository, // transaction repository implementation
	)

	createClientCommand := loadClientCommands.NewCreateClientHandler(storage.ClientRepository, storage.UnitOfWork)
	updateClientCommand := loadClientCommands.NewUpdateClientHandler(storage.ClientRepository, storage.UnitOfWork)
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
