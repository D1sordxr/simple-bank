package app

import (
	storageConfig "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/config"
	GrpcConfig "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/config"
)

type Config struct {
	AppConfig     App                         `yaml:"app"`
	StorageConfig storageConfig.StorageConfig `yaml:"storage"`
	GrpcConfig    GrpcConfig.GrpcConfig       `yaml:"grpc"`
}

type App struct {
	Mode string `yaml:"mode"`
}
