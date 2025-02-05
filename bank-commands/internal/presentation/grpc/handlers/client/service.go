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

// TODO: add event id returning in commands layer

func (s *GrpcService) CreateClient(ctx context.Context, in *services.CreateClientRequest) (*services.CreateClientResponse, error) {
	phones := make([]map[string]int, 0, len(in.Phones))

	for _, data := range in.Phones {
		phone := map[string]int{
			"country": int(data.Country),
			"code":    int(data.Code),
			"number":  int(data.Number),
		}

		phones = append(phones, phone)
	}

	command := commands.CreateClientCommand{
		FirstName:  in.FirstName,
		LastName:   in.LastName,
		MiddleName: in.MiddleName,
		Email:      in.Email,
		Phones:     phones,
	}

	response, err := s.s.CreateClientHandler.Handle(ctx, command)
	if err != nil {
		return nil, err
	}

	return &services.CreateClientResponse{
		Id:      response.ClientID,
		EventID: "nil",
	}, nil
}

// func (s *Service) mustEmbedUnimplementedClientServiceServer() {}
