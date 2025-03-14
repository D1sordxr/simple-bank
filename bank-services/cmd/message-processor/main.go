package main

import (
	consumer2 "github.com/D1sordxr/packages/kafka/consumer"
	pkgProducer "github.com/D1sordxr/packages/kafka/producer"
	pkgLog "github.com/D1sordxr/packages/log"
	pkgPostgres "github.com/D1sordxr/packages/postgres"
	pkgExecutor "github.com/D1sordxr/packages/postgres/executor"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/handlers"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/services"
	loadConfig "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/config-v2"
	loadPostgresProcMsg "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/dao/processed-messages"
	"github.com/D1sordxr/simple-bank/bank-services/internal/presentation"
	"github.com/D1sordxr/simple-bank/bank-services/internal/presentation/consumer"
)

func main() {
	cfg := loadConfig.NewConfig()

	log := pkgLog.Default() // TODO: Default() -> New() (optional)

	pool := pkgPostgres.NewPool(&cfg.Storage)
	executor := pkgExecutor.NewManager(pool)

	txMsgDAO := loadPostgresProcMsg.NewDAO(executor)

	producer := pkgProducer.NewProducer(&cfg.MessageBroker.Producer)

	txMsgProcessorSvc := handlers.NewProcessTransactionHandler(
		log,
		producer,
		cfg.MessageBroker.ProducerTopics.AccountBalanceUpdate,
		txMsgDAO,
		new(services.ProcessDomainSvc),
	)
	// TODO: sagaMsgProcessorSvc :=

	txMsgHandler := consumer.NewHandler(txMsgProcessorSvc)

	txMsgConsumer := consumer2.NewConsumer(
		&cfg.MessageBroker.Consumer,
		cfg.MessageBroker.ConsumerTopics.TransactionCreatedEvent,
		txMsgHandler,
		log,
	)
	sagaMsgConsumer := consumer.NewConsumer(
		&cfg.MessageBroker.Consumer,
	)

	server := consumer.NewServer(
		log,
		txMsgConsumer,
	)
	app := presentation.NewApp(server)
	app.RunApp()
}
