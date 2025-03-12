package handlers

import (
	"context"
	"fmt"
	sharedInterfaces "github.com/D1sordxr/simple-bank/bank-services/internal/application/shared/interfaces"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/interfaces"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/consts"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/account"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/exceptions"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/vo"
)

type ProcessTransactionHandler struct {
	dao      interfaces.ProcessTransactionDAO
	svc      interfaces.ProcessDomainSvc
	producer sharedInterfaces.Producer
}

func NewProcessTransactionHandler(
	dao interfaces.ProcessTransactionDAO,
	producer sharedInterfaces.Producer,
	svc interfaces.ProcessDomainSvc,
) *ProcessTransactionHandler {
	return &ProcessTransactionHandler{
		dao:      dao,
		producer: producer,
		svc:      svc,
	}
}

func (h *ProcessTransactionHandler) Handle(ctx context.Context, dto dto.ProcessDTO) error {
	const op = "Services.TransactionService.ProcessTransaction"
	// TODO: log start

	outboxID, agg, err := h.svc.ParseMessage(dto)
	if err != nil {
		// TODO: log err
		return fmt.Errorf("%s: %w", op, err)
	}

	processed, err := h.dao.IsProcessed(ctx, outboxID)
	if processed {
		// TODO: log msg already processed
		return nil
	}

	messages := make([]account.UpdateEvent, 0, 2)

	switch agg.Type.Value {
	case vo.DepositType: // sends one account update
		messages = append(messages, account.UpdateEvent{
			AccountID:         agg.DestinationAccountID.String(),
			Amount:            agg.Amount.Value,
			BalanceUpdateType: consts.CreditBalanceUpdateType,
		})
	case vo.WithdrawalType: // sends one account update
		messages = append(messages, account.UpdateEvent{
			AccountID:         agg.SourceAccountID.String(),
			Amount:            agg.Amount.Value,
			BalanceUpdateType: consts.DebitBalanceUpdateType,
		})
	case vo.TransferType: // sends two account updates
		messages = append(messages,
			account.UpdateEvent{
				AccountID:         agg.SourceAccountID.String(),
				Amount:            agg.Amount.Value,
				BalanceUpdateType: consts.DebitBalanceUpdateType,
			},
			account.UpdateEvent{
				AccountID:         agg.DestinationAccountID.String(),
				Amount:            agg.Amount.Value,
				BalanceUpdateType: consts.CreditBalanceUpdateType,
			},
		)
	case vo.ReversalType: // TODO: support in main service
		// TODO: log
		return fmt.Errorf("%s: reversal transaction is not supported", op)
	default:
		// TODO: log
		return fmt.Errorf("%s: %w", op, exceptions.InvalidTxType)
	}

	for _, msg := range messages {
		payload, err := h.svc.MarshalMessage(msg)
		if err != nil {
			// TODO: log
			return fmt.Errorf("%s: %w", op, err)
		}

		err = h.producer.SendMessage(ctx, nil, payload)
		if err != nil {
			// TODO: log
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if err = h.dao.SetProcessed(ctx, outboxID); err != nil {
		// TODO: log
		return fmt.Errorf("%s: %w", op, err)
	}

	// TODO: log processed successfully

	return nil
}
