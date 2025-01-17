package app

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	BasicConfigPath = "./configs/app/local.yaml"
)

func NewConfig() *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(BasicConfigPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}
