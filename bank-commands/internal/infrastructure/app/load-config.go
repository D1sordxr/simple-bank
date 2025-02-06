package app

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

const (
	BasicConfigPath = "./configs/app/local.yaml"
)

func NewConfig() *Config {
	var cfg Config

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = BasicConfigPath
	}

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}
