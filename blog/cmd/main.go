package main

import (
	"LearningArch/blog/internal/presentation/di/config"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
	).Run()

}
