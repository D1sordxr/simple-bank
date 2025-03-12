package main

import (
	pkgConsumer "github.com/D1sordxr/packages/kafka/consumer"
	pkgProducer "github.com/D1sordxr/packages/kafka/producer"
	pkgLog "github.com/D1sordxr/packages/log"
	pkgPostgres "github.com/D1sordxr/packages/postgres"
	pkgExecutor "github.com/D1sordxr/packages/postgres/executor"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/handlers"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/services"
	loadConfig "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/config-v2"
	loadPostgresProcMsg "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/dao/processed-messages"
)

func main() {
	// TODO: v0.1.2 packages - add yaml config support
	// TODO: v0.1.2 packages - kafka remove topic creation with new consumer/producer
	// TODO: v0.1.2 packages - kafka add config idempotency support

	cfg := loadConfig.NewConfig()

	log := pkgLog.Default() // TODO: Default() -> New() (optional)

	pool := pkgPostgres.NewPool(&cfg.Storage)
	executor := pkgExecutor.NewManager(pool)

	txMsgDAO := loadPostgresProcMsg.NewDAO(executor)

	// TODO: remove error
	// TODO: change Config -> *Config
	producer, err := pkgProducer.NewProducer(cfg.Producer)
	if err != nil {
		return
	}

	// TODO: packages - producer fixes
	txMsgProcessor := handlers.NewProcessTransactionHandler( // TODO: send topic as arg
		txMsgDAO,
		producer,
		new(services.ProcessDomainSvc),
	)

	// TODO: v0.1.2 packages - read message method for consumer (?)
	txConsumer := pkgConsumer.NewConsumer(&cfg.Consumer) // TODO: send topic as arg

	_, _, _, _ = log, pool, txConsumer, producer

	// TODO: app.Run()
}
