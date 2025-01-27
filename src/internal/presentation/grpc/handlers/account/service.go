package account

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-banking-system/internal/application/account"
	"github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/protobuf/services"
)

type GrpcService struct {
	s *account.Service
	services.UnimplementedAccountServiceServer
}

func NewAccountGrpcService(s *account.Service) *GrpcService {
	return &GrpcService{s: s}
}

func (s *GrpcService) CreateAccount(context.Context, *services.CreateAccountRequest) (*services.CreateAccountResponse, error) {

	// TODO: ...

	return nil, fmt.Errorf("not implemented")
}
