package app

import (
	storageConfig "github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/config"
	GrpcConfig "github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/config"
)

type Config struct {
	AppConfig     App                         `yaml:"app"`
	StorageConfig storageConfig.StorageConfig `yaml:"storage"`
	GrpcConfig    GrpcConfig.GrpcConfig       `yaml:"grpc"`
}

type App struct {
	Mode string `yaml:"mode"`
}
