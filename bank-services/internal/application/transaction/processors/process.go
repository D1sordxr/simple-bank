package processors

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
	log           sharedInterfaces.Logger
	producer      sharedInterfaces.Producer
	producerTopic string
	dao           interfaces.ProcessTransactionDAO
	svc           interfaces.ProcessDomainSvc
}

func NewProcessTransactionHandler(
	log sharedInterfaces.Logger,
	producer sharedInterfaces.Producer,
	topic string,
	dao interfaces.ProcessTransactionDAO,
	svc interfaces.ProcessDomainSvc,
) *ProcessTransactionHandler {
	return &ProcessTransactionHandler{
		log:           log,
		dao:           dao,
		producerTopic: topic,
		producer:      producer,
		svc:           svc,
	}
}

func (h *ProcessTransactionHandler) Process(ctx context.Context, dto dto.ProcessDTO) error {
	const op = "Services.TransactionService.ProcessTransaction"

	h.log.Info("Attempting to process message...")

	outboxID, agg, err := h.svc.ParseMessage(dto)
	if err != nil {
		h.log.Errorw("Error parsing message", "error", err.Error())
		return fmt.Errorf("%s: %w", op, err)
	}

	processed, err := h.dao.IsProcessed(ctx, outboxID)
	switch {
	case err != nil:
		h.log.Errorw("Error checking message status",
			"error", err.Error(),
			"outboxID", outboxID,
		)
		return fmt.Errorf("%s: %w", op, err)
	case processed:
		h.log.Infow("Message already processed", "outboxID", outboxID)
		return nil
	}

	h.log.Infow("Start processing message", "outboxID", outboxID)

	messages := make(account.UpdateEvents, 0, 2)

	switch agg.Type.Value {
	case vo.DepositType: // sends one account update event
		messages = append(messages, account.UpdateEvent{
			AccountID:         agg.DestinationAccountID.String(),
			Amount:            agg.Amount.Value,
			BalanceUpdateType: consts.CreditBalanceUpdateType,
			TransactionID:     agg.TransactionID.String(),
		})
	case vo.WithdrawalType: // sends one account update event
		messages = append(messages, account.UpdateEvent{
			AccountID:         agg.SourceAccountID.String(),
			Amount:            agg.Amount.Value,
			BalanceUpdateType: consts.DebitBalanceUpdateType,
			TransactionID:     agg.TransactionID.String(),
		})
	case vo.TransferType: // sends two account updates event
		messages = append(messages,
			account.UpdateEvent{
				AccountID:         agg.SourceAccountID.String(),
				Amount:            agg.Amount.Value,
				BalanceUpdateType: consts.DebitBalanceUpdateType,
				TransactionID:     agg.TransactionID.String(),
			},
			account.UpdateEvent{
				AccountID:         agg.DestinationAccountID.String(),
				Amount:            agg.Amount.Value,
				BalanceUpdateType: consts.CreditBalanceUpdateType,
				TransactionID:     agg.TransactionID.String(),
			},
		)
	case vo.ReversalType: // TODO: support in main service
		// TODO: log
		return fmt.Errorf("%s: reversal transaction is not supported", op)
	default:
		h.log.Errorw("Invalid transaction type",
			"type", agg.Type.Value,
			"outboxID", outboxID,
		)
		return fmt.Errorf("%s: %w", op, exceptions.InvalidTxType)
	}

	payload, err := h.svc.MarshalMessage(messages)
	if err != nil {
		h.log.Errorw("Failed to marshal message",
			"error", err.Error(),
			"outboxID", outboxID,
		)
		return fmt.Errorf("%s: %w", op, err)
	}

	err = h.producer.SendMessage(ctx, h.producerTopic, nil, payload)
	if err != nil {
		h.log.Errorw("Failed to send message",
			"error", err.Error(),
			"outboxID", outboxID,
		)
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = h.dao.SetProcessed(ctx, outboxID); err != nil {
		h.log.Errorw("Error setting new message status",
			"error", err.Error(),
			"outboxID", outboxID,
		)
		return fmt.Errorf("%s: %w", op, err)
	}

	h.log.Infow("Message processed successfully", "outboxID", outboxID)

	return nil
}
