package config

import (
	db "LearningArch/blog/internal/infrastructure/db/config"
	api "LearningArch/blog/internal/presentation/api/config"
)

type AppConfig struct {
	Mode string
}

type Config struct {
	AppConfig     `toml:"app"`
	api.APIConfig `toml:"api"`
	db.DBConfig   `toml:"db"`
}
