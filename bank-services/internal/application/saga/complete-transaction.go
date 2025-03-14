package saga

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/saga/interfaces"
	sharedInterfaces "github.com/D1sordxr/simple-bank/bank-services/internal/application/shared/interfaces"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	pbServices "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/protobuf/services"
	"google.golang.org/grpc"
	"log"
	"time"
)

type CompleteTransactionProcessor struct {
	log               sharedInterfaces.Logger
	svc               interfaces.CompleteTransactionDomainSvc
	accountClient     pbServices.AccountServiceClient
	transactionClient pbServices.AccountServiceClient
}

func NewCompleteTransactionProcessor(
	accountConn *grpc.ClientConn,
	transactionConn *grpc.ClientConn,
) *CompleteTransactionProcessor {
	return &CompleteTransactionProcessor{
		//accountClient:     pbServices.NewAccountServiceClient(accountConn),
		//transactionClient: pbServices.NewTransactionServiceClient(transactionConn),
	}
}

func (p *CompleteTransactionProcessor) Process(ctx context.Context, dto dto.ProcessDTO) error {
	const op = "application.saga.CompleteTransactionProcessor.Process"

	p.log.Info("Starting Saga...")

	accountUpdates, err := p.svc.UnmarshalData(dto)
	if err != nil {
		// TODO: log
		return fmt.Errorf("%s: %w", op, err)
	}
	for _, update := range accountUpdates {
		request :=
	}

	// Шаг 1: Обновление аккаунта
	if err := s.updateAccount(ctx, dto.AccountID); err != nil {
		return fmt.Errorf("failed to update account: %w", err)
	}

	// Шаг 2: Обновление транзакции
	if err := s.updateTransaction(ctx, dto.TransactionID); err != nil {
		// Откат (компенсирующая операция)
		log.Println("Rolling back account update...")
		s.rollbackAccount(ctx, dto.AccountID)
		return fmt.Errorf("failed to update transaction: %w", err)
	}

	log.Println("Saga completed successfully")
	return nil
}

func (s *Saga) updateAccount(ctx context.Context, accountID string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := s.accountClient.UpdateAccountHandler(ctx, &pb.UpdateAccountRequest{AccountId: accountID})
	if err != nil {
		log.Printf("Error updating account: %v", err)
	}
	return err
}

func (s *Saga) updateTransaction(ctx context.Context, transactionID string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := s.transactionClient.UpdateTransactionHandler(ctx, &pb.UpdateTransactionRequest{TransactionId: transactionID})
	if err != nil {
		log.Printf("Error updating transaction: %v", err)
	}
	return err
}

func (s *Saga) rollbackAccount(ctx context.Context, accountID string) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := s.accountClient.RollbackAccountHandler(ctx, &pb.RollbackAccountRequest{AccountId: accountID})
	if err != nil {
		log.Printf("Error rolling back account: %v", err)
	}
}
