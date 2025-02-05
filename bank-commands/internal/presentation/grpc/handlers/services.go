package handlers

import (
	"github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/protobuf/services"
)

type Services struct {
	ClientService      services.ClientServiceServer
	AccountService     services.AccountServiceServer
	TransactionService services.TransactionServiceServer
}

func NewGrpcServices(
	cs services.ClientServiceServer,
	as services.AccountServiceServer,
	ts services.TransactionServiceServer,
) *Services {
	return &Services{
		ClientService:      cs,
		AccountService:     as,
		TransactionService: ts,
	}
}
