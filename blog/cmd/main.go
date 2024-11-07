package main

import (
	"LearningArch/blog/internal/presentation/di/config"
	"fmt"
	"go.uber.org/fx"
)

func main() {
	fmt.Println(config.Module.String())
	fx.New(
		config.Module,
	).Run()

}
