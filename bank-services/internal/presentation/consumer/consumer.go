package consumer

import (
	"context"
	"github.com/D1sordxr/packages/kafka/consumer"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Processor interface {
	Process(ctx context.Context, msg kafka.Message) error
}

type Consumer struct {
	proc   Processor
	reader *kafka.Reader
}

func NewConsumer(config *consumer.Config, proc Processor) *Consumer {
	return &Consumer{
		reader: consumer.NewConsumer(config),
		proc:   proc,
	}
}

func (c *Consumer) ReadAndProcess(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			// log.Println("Consumer: получен сигнал завершения, останавливаем обработку...")
			return
		default:
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				// log.Printf("Ошибка при чтении сообщения: %v", err)
				continue
			}

			if err = c.proc.Process(ctx, msg); err != nil {
				// log.Printf("Ошибка при обработке сообщения: %v", err)
			}
		}
	}
}

type Server struct {
	consumers []*Consumer
}

func NewServer(consumers ...*Consumer) *Server {
	return &Server{consumers: consumers}
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

		// log.Println("Stop signal received, shutting down...")

		for _, r := range s.consumers {
			_ = r.reader.Close()
		}
		cancel()
	}()

	for _, c := range s.consumers {
		wg.Add(1)
		go c.ReadAndProcess(ctx, wg)
	}

	select {
	case <-ctx.Done():
		// log.Println("Context cancelled, shutting down all processes...")
	case _ = <-errorsChannel:
		// log.Printf("Application encountered an error: %v", err)
		// TODO: cancel() for critical errors
	}

	wg.Wait()
	// log.Println("Server gracefully stopped")
}
