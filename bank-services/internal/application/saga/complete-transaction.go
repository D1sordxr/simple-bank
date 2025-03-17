package saga

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/saga/interfaces"
	sharedInterfaces "github.com/D1sordxr/simple-bank/bank-services/internal/application/shared/interfaces"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/vo"
	pbServices "github.com/D1sordxr/simple-bank/bank-services/internal/presentation/grpc/protobuf/services"
	"google.golang.org/grpc"
)

type CompleteTransactionSaga struct {
	log               sharedInterfaces.Logger
	svc               interfaces.CompleteTransactionDomainSvc
	accountClient     pbServices.AccountServiceClient
	transactionClient pbServices.TransactionServiceClient
	// TODO: rollbackClient pbServices.RollbackServiceClient
}

func NewCompleteTransactionProcessor(
	accountConn *grpc.ClientConn,
	transactionConn *grpc.ClientConn,
) *CompleteTransactionSaga {
	return &CompleteTransactionSaga{
		//accountClient:     pbServices.NewAccountServiceClient(accountConn),
		//transactionClient: pbServices.NewTransactionServiceClient(transactionConn),
	}
}

func (s *CompleteTransactionSaga) Process(ctx context.Context, dto dto.ProcessDTO) error {
	const op = "application.saga.CompleteTransactionProcessor.Process"

	s.log.Info("Starting complete transaction saga...")

	txID, accountUpdates, err := s.svc.UnmarshalData(dto)
	if err != nil {
		s.log.Errorw("failed to unmarshal data", "error", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	var rollbackEvents = make(event.RollbackEvents, 0, len(accountUpdates)+1) // +1 for transaction update
	defer s.rollback(rollbackEvents, &err)

	for _, update := range accountUpdates {
		request := &pbServices.UpdateAccountRequest{
			AccountID:   update.AccountID,
			Amount:      float32(update.Amount),
			BalanceType: update.BalanceUpdateType,
		}
		response, err := s.accountClient.UpdateAccount(ctx, request)
		if err != nil {
			s.log.Errorw("failed to update account", "error", err,
				"account_id", update.AccountID,
				"transaction_id", txID,
			)
			return fmt.Errorf("%s: %w", op, err)
		}
		rollbackEvents = append(rollbackEvents, event.Rollback{EventID: response.EventID}) // TODO: rollback events
	}

	txRequest := &pbServices.UpdateTransactionRequest{
		TransactionID: txID,
		Status:        vo.StatusCompleted,
		FailureReason: "",
	}
	txResponse, err := s.transactionClient.UpdateTransaction(ctx, txRequest)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	rollbackEvents = append(rollbackEvents, event.Rollback{EventID: txResponse.EventID}) // TODO: rollback events

	s.log.Info("Saga completed successfully")

	return nil
}

func (s *CompleteTransactionSaga) rollback(events event.RollbackEvents, err *error) {
	const op = "application.saga.CompleteTransactionSaga.Process.rollback"

	s.log.Info("Starting rollback transaction saga...")

	_ = context.Background() // or with timeout

	// TODO: rollback events handler

	s.log.Info("Saga rollback completed successfully")
}
