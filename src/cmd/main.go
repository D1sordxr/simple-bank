package main

import (
	loadStorage "github.com/D1sordxr/simple-banking-system/internal/infrastructure"
	loadConfig "github.com/D1sordxr/simple-banking-system/internal/infrastructure/app"
	loadPostgresConnection "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres"
	loadPostgresAccountRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/account"
	loadPostgresClientRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/client"
	loadPostgresTransactionRepo "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/repositories/transaction"
	loadPostgresUoW "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/uow"
)

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
	clientRepo := loadPostgresClientRepo.NewClientRepository(databaseConn)
	accountRepo := loadPostgresAccountRepo.NewAccountRepository(databaseConn)
	transactionRepo := loadPostgresTransactionRepo.NewTransactionRepository(databaseConn)

	storage := loadStorage.NewStorage( // TODO: finish client and account repos
		uowManager,      // uow implementation
		clientRepo,      // client repository implementation
		accountRepo,     // account repository implementation
		transactionRepo, // transaction repository implementation
	)
}
