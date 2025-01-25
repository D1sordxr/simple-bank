package grpc

import "github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/protobuf/services"

type Services struct {
	cs services.ClientServiceServer
	// as *pkg.AccountServiceClient
	// ts *pkg.TransactionServiceClient
}

func NewGrpcServices(
	cs services.ClientServiceServer,
	// as *pkg.AccountServiceClient,
	// ts *pkg.TransactionServiceClient,
) *Services {
	return &Services{
		cs: cs,
	}
}
