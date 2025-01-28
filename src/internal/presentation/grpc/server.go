package grpc

import (
	"fmt"
	"github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/config"
	"github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/handlers"
	pbServices "github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/protobuf/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	Config       *config.GrpcConfig
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
		Config:       config,
		GrpcServices: services,
		Server:       server,
	}
}

func (s *Server) RegisterServices() {
	pbServices.RegisterClientServiceServer(s.Server, s.GrpcServices.ClientService)
	pbServices.RegisterAccountServiceServer(s.Server, s.GrpcServices.AccountService)
	pbServices.RegisterTransactionServiceServer(s.Server, s.GrpcServices.TransactionService)

	reflection.Register(s.Server)
}

func (s *Server) Run() {
	s.RegisterServices()

	// TODO: RunGrpc in goroutine and listen for stop call

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Config.Port))

	if err != nil {
		// TODO: fatal
		return
	}

	if err = s.Server.Serve(l); err != nil {
		// TODO: fatal
		return
	}

	return
	// TODO: Listen and serve

}

func (s *Server) Down() {
	s.Server.GracefulStop()
}
