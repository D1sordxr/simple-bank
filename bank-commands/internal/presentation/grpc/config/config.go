package config

import "time"

type GrpcConfig struct {
	Port              int           `yaml:"port"`
	Timeout           time.Duration `yaml:"timeout"`             // time.Second
	Time              time.Duration `yaml:"time"`                // time.Minute
	MaxConnectionIdle time.Duration `yaml:"max_connection_idle"` // time.Minute
	MaxConnectionAge  time.Duration `yaml:"max_connection_age"`  // time.Minute
}
