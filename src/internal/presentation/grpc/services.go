package grpc

import (
	"github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/protobuf/services"
)

type Services struct {
	cs services.ClientServiceServer
	as services.AccountServiceServer
	ts services.TransactionServiceServer
}

func NewGrpcServices(
	cs services.ClientServiceServer,
	as services.AccountServiceServer,
	ts services.TransactionServiceServer,
) *Services {
	return &Services{
		cs: cs,
		as: as,
		ts: ts,
	}
}
