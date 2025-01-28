package grpc

import (
	"fmt"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/app/logger"
	"github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/config"
	"github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/handlers"
	pbServices "github.com/D1sordxr/simple-banking-system/internal/presentation/grpc/protobuf/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	Config       *config.GrpcConfig
	Logger       *logger.Logger
	GrpcServices *handlers.Services
	Server       *grpc.Server
}

func NewGrpcServer(config *config.GrpcConfig,
	logger *logger.Logger,
	services *handlers.Services) *Server {
	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:              config.Time,
		Timeout:           config.Timeout,
		MaxConnectionIdle: config.MaxConnectionIdle,
		MaxConnectionAge:  config.MaxConnectionAge,
	}))

	return &Server{
		Logger:       logger,
		Config:       config,
		GrpcServices: services,
		Server:       server,
	}
}

func (s *Server) Run() {
	var err error
	errorsChannel := make(chan error, 1)

	go func() {
		if err = s.RunGrpcServer(); err != nil {
			errorsChannel <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-stop:
		s.Logger.Info("Stopping application...", s.Logger.String("signal", "stop"))
	case err = <-errorsChannel:
		s.Logger.Error("Application encountered an error", s.Logger.String("error", err.Error()))
	}

	s.Down()
	s.Logger.Info("Gracefully stopped")
}

func (s *Server) RunGrpcServer() error {
	const op = "grpcServer.Run"
	port := s.Config.Port

	s.registerServices()
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.Server.Serve(l)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.Logger.Info("GRPC server successfully started")
	return nil
}

func (s *Server) Down() {
	s.Server.GracefulStop()
}

func (s *Server) registerServices() {
	pbServices.RegisterClientServiceServer(s.Server, s.GrpcServices.ClientService)
	pbServices.RegisterAccountServiceServer(s.Server, s.GrpcServices.AccountService)
	pbServices.RegisterTransactionServiceServer(s.Server, s.GrpcServices.TransactionService)

	reflection.Register(s.Server)
}
