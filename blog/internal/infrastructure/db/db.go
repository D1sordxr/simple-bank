package db

import (
	"LearningArch/blog/internal/presentation/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func BuildConnection(config config.Config) *gorm.DB {
	gormConfig := gorm.Config{}
	if !config.Logging {
		gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	} // TODO: else { gormConfig.Logger = logger.GetGormLogger() }

	db, err := gorm.Open(postgres.Open(config.ConnectionString()), &gormConfig)
	if err != nil {
		panic(err.Error())
	}

	// TODO: if config.Migration { migrate(db) }

	return db
}
