package main

import (
	"LearningArch/bank-workers/internal/infrastructure/app"
	consumer2 "github.com/D1sordxr/packages/kafka/consumer"
	logPkg "github.com/D1sordxr/packages/log"
	"github.com/D1sordxr/packages/postgres"
)

func main() {
	// TODO: v0.1.2 packages - add yaml config support

	cfg := app.NewConfig()

	log := logPkg.Default() // TODO: Default() -> New() (optional)

	pool := postgres.NewPool(&cfg.Storage)

	// TODO: v0.1.2 packages - read message method for consumer (!)
	consumer := consumer2.NewConsumer(&cfg.Consumer)

	_, _, _ = log, pool, consumer

	// TODO: mediator

	// TODO: app.Run()
}
