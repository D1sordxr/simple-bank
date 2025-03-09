package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/shared/interfaces"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/consts"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/account"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/exceptions"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/services"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/vo"
)

type ProcessTransactionHandler struct {
	producer interfaces.Producer
	svc      services.IProcessDomainSvc
}

func NewProcessTransactionHandler(
	producer interfaces.Producer,
	svc services.IProcessDomainSvc,
) *ProcessTransactionHandler {
	return &ProcessTransactionHandler{
		producer: producer,
		svc:      svc,
	}
}

func (h *ProcessTransactionHandler) Handle(ctx context.Context, dto dto.ProcessDTO) error {
	const op = "Services.TransactionService.ProcessTransaction"
	// log start

	agg, err := h.svc.ParseMessage(dto.ByteData)
	if err != nil {
		// log
		return fmt.Errorf("%s: %w", op, err)
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
		return fmt.Errorf("%s: ReversalType is not implemented", op)
	default:
		return fmt.Errorf("%s: %w", op, exceptions.InvalidTxType)
	}

	for _, msg := range messages {
		payload, err := json.Marshal(msg)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		err = h.producer.SendMessage(ctx, nil, payload)
		if err != nil {
			return fmt.Errorf("%s: я%w", op, err)
		}
	}

	return nil
}
