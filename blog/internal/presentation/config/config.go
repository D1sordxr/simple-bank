package config

import (
	api "LearningArch/blog/internal/infrastructure/api/config"
	db "LearningArch/blog/internal/infrastructure/db/config"
	logger "LearningArch/blog/internal/infrastructure/logger/config"
)

type AppConfig struct {
	Mode string
}

type Config struct {
	AppConfig           `toml:"app"`
	api.APIConfig       `toml:"api"`
	db.DBConfig         `toml:"db"`
	logger.LoggerConfig `toml:"logger"`
}
