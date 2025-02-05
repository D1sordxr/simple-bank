package account

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/account"
	"github.com/D1sordxr/simple-banking-system/internal/application/account/commands"
	"github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/protobuf/services"
)

type GrpcService struct {
	s *account.Service
	services.UnimplementedAccountServiceServer
}

func NewAccountGrpcService(s *account.Service) *GrpcService {
	return &GrpcService{s: s}
}

// TODO: add event id returning in commands layer

func (s *GrpcService) CreateAccount(ctx context.Context, req *services.CreateAccountRequest) (*services.CreateAccountResponse, error) {
	command := commands.CreateAccountCommand{
		ClientID: req.ClientID,
		Currency: req.Currency,
	}

	response, err := s.s.CreateAccountHandler.Handle(ctx, command)
	if err != nil {
		return nil, err
	}

	return &services.CreateAccountResponse{
		Id:      response.AccountID,
		EventID: "nil",
	}, nil
}
