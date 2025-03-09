package main

import (
	pkgConsumer "github.com/D1sordxr/packages/kafka/consumer"
	pkgProducer "github.com/D1sordxr/packages/kafka/producer"
	pkgLog "github.com/D1sordxr/packages/log"
	pkgPostgres "github.com/D1sordxr/packages/postgres"
	loadConfig "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/config-v2"
)

func main() {
	// TODO: v0.1.2 packages - add yaml config support
	// TODO: v0.1.2 packages - remove topic creation with new consumer/producer

	cfg := loadConfig.NewConfig()

	log := pkgLog.Default() // TODO: Default() -> New() (optional)

	pool := pkgPostgres.NewPool(&cfg.Storage)

	// TODO: remove error
	// TODO: change Config -> *Config
	producer, err := pkgProducer.NewProducer(cfg.Producer)
	if err != nil {
		return
	}

	// TODO: v0.1.2 packages - read message method for consumer (?)
	consumer := pkgConsumer.NewConsumer(&cfg.Consumer)

	_, _, _, _ = log, pool, consumer, producer

	// TODO: app.Run()
}
