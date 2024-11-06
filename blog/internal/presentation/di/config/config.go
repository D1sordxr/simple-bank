package config

import (
	"LearningArch/blog/internal/presentation/config"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"presentation.config",
	fx.Provide(
		config.NewConfig,
		config.NewDBConfig,
		config.NewAPIConfig,
		config.NewLoggerConfig,
		config.NewAppConfig,
	),
)
