package main

import (
	"LearningArch/blog/internal/infrastructure/di"
	"LearningArch/blog/internal/presentation/di/config"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		di.Module,
	).Run()

}
