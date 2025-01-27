package transaction

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-banking-system/internal/application/transaction"
	"github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/protobuf/services"
)

type GrpcService struct {
	s *transaction.Service
	services.UnimplementedTransactionServiceServer
}

func NewTransactionGrpcService(s *transaction.Service) *GrpcService {
	return &GrpcService{s: s}
}

func (s *GrpcService) CreateTransaction(context.Context, *services.CreateTransactionRequest) (*services.CreateTransactionResponse, error) {

	// TODO: ...

	return nil, fmt.Errorf("not implemented")
}
