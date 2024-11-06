package config

import (
	api "LearningArch/blog/internal/infrastructure/api/config"
	load "LearningArch/blog/internal/infrastructure/config"
	db "LearningArch/blog/internal/infrastructure/db/config"
	logger "LearningArch/blog/internal/infrastructure/logger/config"
)

func NewConfig() Config {
	var config Config
	load.LoadConfig(&config, "", "")
	return config
}

func NewAppConfig(config Config) AppConfig {
	return config.AppConfig
}

func NewDBConfig(config Config) db.DBConfig {
	return config.DBConfig
}

func NewAPIConfig(config Config) api.APIConfig {
	config.APIConfig.Mode = config.AppConfig.Mode
	return config.APIConfig
}

func NewLoggerConfig(config Config) logger.LoggerConfig {
	config.LoggerConfig.Mode = config.AppConfig.Mode
	return config.LoggerConfig
}
