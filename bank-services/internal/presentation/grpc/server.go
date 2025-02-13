package grpc

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/logger"
	"github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/config"
	"github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/handlers"
	pbServices "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/protobuf/services"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err = s.RunGrpcServer(); err != nil {
			errorsChannel <- err
		}
	}()

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
		<-stop
		s.Logger.Info("Stop signal received...")
		cancel()
	}()

	select {
	case <-ctx.Done():
		s.Logger.Info("Stopping application...", s.Logger.String("reason", "stop signal"))
	case err = <-errorsChannel:
		s.Logger.Error("Application encountered an error", s.Logger.String("error", err.Error()))
	}

	s.Down()
	s.Logger.Info("Gracefully stopped")
}

func (s *Server) RunGrpcServer() error {
	const op = "grpcServer.Run"

	s.registerServices()
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Config.Port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.Logger.Info("GRPC server successfully started", s.Logger.String("port", fmt.Sprintf(":%d", s.Config.Port)))

	err = s.Server.Serve(l)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Server) Down() {
	s.Logger.Info("Shutting down GRPC server...")
	s.Server.GracefulStop()
	s.Logger.Info("GRPC server stopped successfully")
}

func (s *Server) registerServices() {
	pbServices.RegisterClientServiceServer(s.Server, s.GrpcServices.ClientService)
	pbServices.RegisterAccountServiceServer(s.Server, s.GrpcServices.AccountService)
	pbServices.RegisterTransactionServiceServer(s.Server, s.GrpcServices.TransactionService)

	reflection.Register(s.Server)
	s.Logger.Info("gRPC services registered successfully")
}
