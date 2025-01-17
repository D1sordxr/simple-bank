package app

import (
	storageConfig "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/config"
	grpcConfig "github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/config"
)

type Config struct {
	AppConfig     App                         `yaml:"app"`
	StorageConfig storageConfig.StorageConfig `yaml:"storage"`
	GRPCConfig    grpcConfig.GRPCConfig       `yaml:"grpc"`
}

type App struct {
	Mode string `yaml:"mode"`
}
