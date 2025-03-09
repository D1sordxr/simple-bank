package handlers

import (
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/exceptions"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/services"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/vo"
)

type ProcessTransactionHandler struct {
	svc services.IProcessDomainSvc
}

func NewProcessTransactionHandler(
	svc services.IProcessDomainSvc,
) *ProcessTransactionHandler {
	return &ProcessTransactionHandler{
		svc: svc,
	}
}

func (h *ProcessTransactionHandler) Handle(dto dto.ProcessDTO) error {
	const op = ""

	agg, err := h.svc.ParseMessage(dto.ByteData)
	if err != nil {
		// log
		return fmt.Errorf("%s: %w", op, err)
	}

	var messages [][]byte

	switch agg.Type.Value {
	case vo.TransferType:
		//
	case vo.DepositType:
		//
	case vo.WithdrawalType:
		//
	case vo.ReversalType:
		//
	default:
		return fmt.Errorf("%s: %w", op, exceptions.InvalidTxType)
	}

	for range messages {
		// msg send to kafka and add sender for UpdateAccountHandler
	}

	return nil
}
