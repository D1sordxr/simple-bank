package db

import (
	"LearningArch/blog/internal/infrastructure/db"
	"LearningArch/blog/internal/infrastructure/db/dao"
	base "LearningArch/blog/internal/infrastructure/db/repo"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func NewBaseRepo(connection *gorm.DB) base.BaseGormRepo {
	return base.BaseGormRepo{Session: connection}
}

func NewBaseDAO(connection *gorm.DB) dao.BaseGormDAO {
	return dao.BaseGormDAO{Session: connection}
}

var Module = fx.Module(
	"infrastructure.di",
	fx.Provide(
		db.BuildConnection,
		NewBaseRepo,
		NewBaseDAO,
	),
)
