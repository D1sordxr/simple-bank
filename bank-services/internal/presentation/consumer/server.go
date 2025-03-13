package consumer

import (
	"context"
	"fmt"
	"github.com/D1sordxr/packages/kafka/consumer"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type log interface {
	Info(msg string)
	Error(msg string)
}

type Server struct {
	log       log
	consumers []*consumer.Consumer
}

func NewServer(log log, consumers ...*consumer.Consumer) *Server {
	return &Server{
		log:       log,
		consumers: consumers,
	}
}

func (s *Server) Run() {
	wg := &sync.WaitGroup{}
	errorsChannel := make(chan error, 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
		<-stop

		s.log.Info("Stop signal received, shutting down...")

		for _, r := range s.consumers {
			_ = r.Reader.Close()
		}
		cancel()
	}()

	for _, c := range s.consumers {
		wg.Add(1)
		go c.Consume(ctx, wg)
	}

	select {
	case <-ctx.Done():
		s.log.Info("Context cancelled, shutting down all processes...")
	case err := <-errorsChannel:
		s.log.Error(fmt.Sprintf("Application encountered an error: %v", err.Error()))
		// TODO: cancel() for critical errors
	}

	wg.Wait()
	s.log.Info("Server gracefully stopped")
}
