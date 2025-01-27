package client

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/client"
	"github.com/D1sordxr/simple-banking-system/internal/application/client/commands"
	"github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/protobuf/services"
)

type GrpcService struct {
	s *client.Service
	services.UnimplementedClientServiceServer
}

func NewClientGrpcService(s *client.Service) *GrpcService {
	return &GrpcService{s: s}
}

func (s *GrpcService) CreateClient(ctx context.Context, in *services.CreateClientRequest) (*services.CreateClientResponse, error) {

	// TODO: ...

	var command commands.CreateClientCommand

	response, err := s.s.CreateClientHandler.Handle(ctx, command)
	if err != nil {
		return nil, err
	}
	
	return &services.CreateClientResponse{
		Id: response.ClientID,
	}, nil
}

// func (s *Service) mustEmbedUnimplementedClientServiceServer() {}
