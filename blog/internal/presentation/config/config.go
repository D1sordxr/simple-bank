package config

type AppConfig struct {
	Mode string
}

type Config struct {
	AppConfig `toml:"app"`
}
