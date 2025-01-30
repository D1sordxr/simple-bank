package transaction

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/application/transaction"
	"github.com/D1sordxr/simple-banking-system/internal/application/transaction/commands"
	"github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/protobuf/services"
)

type GrpcService struct {
	s *transaction.Service
	services.UnimplementedTransactionServiceServer
}

func NewTransactionGrpcService(s *transaction.Service) *GrpcService {
	return &GrpcService{s: s}
}

// TODO: add event id returning in commands layer

func (s *GrpcService) CreateTransaction(ctx context.Context, req *services.CreateTransactionRequest) (*services.CreateTransactionResponse, error) {
	command := commands.CreateTransactionCommand{
		SourceAccountID:      req.SourceAccountID,
		DestinationAccountID: req.DestinationAccountID,
		Currency:             req.Currency,
		Amount:               float64(req.Amount),
		Type:                 req.Type,
		Description:          req.Description,
	}

	response, err := s.s.CreateTransactionHandler.Handle(ctx, command)
	if err != nil {
		return nil, err
	}

	return &services.CreateTransactionResponse{
		Id:      response.TransactionID,
		EventID: "nil",
	}, nil
}
