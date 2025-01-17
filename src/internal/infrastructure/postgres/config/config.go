package config

import "fmt"

type StorageConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Database  string `yaml:"database"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Migration bool   `yaml:"migration"`
}

// ConnectionString returns DSN
func (c *StorageConfig) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		c.Host, c.User, c.Password, c.Database, c.Port,
	)
}
