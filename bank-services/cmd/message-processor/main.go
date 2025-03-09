package main

import (
	pkgConsumer "github.com/D1sordxr/packages/kafka/consumer"
	pkgLog "github.com/D1sordxr/packages/log"
	pkgPostgres "github.com/D1sordxr/packages/postgres"
	loadConfig "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/config-v2"
)

func main() {
	// TODO: Monorepo in bank-services for sharing domain

	// TODO: v0.1.2 packages - add yaml config support

	cfg := loadConfig.NewConfig()

	log := pkgLog.Default() // TODO: Default() -> New() (optional)

	pool := pkgPostgres.NewPool(&cfg.Storage)

	// TODO: v0.1.2 packages - read message method for consumer (!)
	consumer := pkgConsumer.NewConsumer(&cfg.Consumer)

	_, _, _ = log, pool, consumer

	// TODO: mediator

	// TODO: app.Run()
}
