package grpc

import (
	"github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/config"
	"github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/handlers"
	pbServices "github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/protobuf/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Server struct {
	GrpcServices *handlers.Services
	Server       *grpc.Server
}

func NewGrpcServer(config *config.GrpcConfig, services *handlers.Services) *Server {
	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:              config.Time,
		Timeout:           config.Timeout,
		MaxConnectionIdle: config.MaxConnectionIdle,
		MaxConnectionAge:  config.MaxConnectionAge,
	}))

	return &Server{
		GrpcServices: services,
		Server:       server,
	}
}

func (s *Server) RegisterServices() {
	pbServices.RegisterClientServiceServer(s.Server, s.GrpcServices.ClientService)
	pbServices.RegisterAccountServiceServer(s.Server, s.GrpcServices.AccountService)
	pbServices.RegisterTransactionServiceServer(s.Server, s.GrpcServices.TransactionService)
}

func (s *Server) Run() {
	s.RegisterServices()

	// TODO: Listen and serve

}

func (s *Server) Down() {

	// TODO: Graceful shutdown

}
