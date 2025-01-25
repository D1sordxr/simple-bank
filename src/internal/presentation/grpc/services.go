package grpc

import services "github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/pb"

type Services struct {
	cs *services.ClientServiceClient
	// as *pkg.AccountServiceClient
	// ts *pkg.TransactionServiceClient
}

func NewGRPCServices(
	cs *services.ClientServiceClient,
	// as *pkg.AccountServiceClient,
	// ts *pkg.TransactionServiceClient,
) *Services {
	return &Services{
		cs: cs,
	}
}
