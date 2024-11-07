package di

import (
	"LearningArch/blog/internal/infrastructure/di/factories/db"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"infrastructure.di",
	fx.Options(
		db.Module,
	),
)
