package main

import (
	consumer2 "github.com/D1sordxr/packages/kafka/consumer"
	pkgLog "github.com/D1sordxr/packages/log"
	pkgPostgres "github.com/D1sordxr/packages/postgres"
	pkgExecutor "github.com/D1sordxr/packages/postgres/executor"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/handlers"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/services"
	loadConfig "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/config-v2"
	loadPostgresProcMsg "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/dao/processed-messages"
	"github.com/D1sordxr/simple-bank/bank-services/internal/presentation"
	"github.com/D1sordxr/simple-bank/bank-services/internal/presentation/consumer"
	"github.com/D1sordxr/simple-bank/bank-services/internal/presentation/consumer/handlers/transaction"
)

func main() {
	cfg := loadConfig.NewConfig()

	log := pkgLog.Default() // TODO: Default() -> New() (optional)

	pool := pkgPostgres.NewPool(&cfg.Storage)
	executor := pkgExecutor.NewManager(pool)

	txMsgDAO := loadPostgresProcMsg.NewDAO(executor)

	//producer, err := pkgProducer.NewProducer(cfg.Producer)
	//if err != nil {
	//	return
	//}

	txMsgProcessorSvc := handlers.NewProcessTransactionHandler(
		txMsgDAO,
		producer,
		new(services.ProcessDomainSvc),
	)

	txMsgHandler := transaction.NewHandler(txMsgProcessorSvc)

	txMsgConsumer := consumer2.NewConsumer(
		&cfg.Consumer,
		cfg.ConsumerTopics.Transaction,
		txMsgHandler,
		log,
	)

	server := consumer.NewServer(
		log,
		txMsgConsumer,
	)
	app := presentation.NewApp(server)
	app.RunApp()
}
