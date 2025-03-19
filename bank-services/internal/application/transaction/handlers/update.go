package handlers

import (
	"context"
	"fmt"
	sharedInterfaces "github.com/D1sordxr/simple-bank/bank-services/internal/application/shared/interfaces"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/commands"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/interfaces"
)

type UpdateTransactionHandler struct {
	log        sharedInterfaces.Logger
	uow        sharedInterfaces.IUnitOfWork
	eventRepo  sharedInterfaces.EventRepo
	outboxRepo sharedInterfaces.OutboxRepo
	svc        interfaces.UpdateDomainSvc
}

func NewUpdateTransactionHandler(
	log sharedInterfaces.Logger,
	uow sharedInterfaces.IUnitOfWork,
	eventRepo sharedInterfaces.EventRepo,
	outboxRepo sharedInterfaces.OutboxRepo,
	svc interfaces.UpdateDomainSvc,
) *UpdateTransactionHandler {
	return &UpdateTransactionHandler{
		log:        log,
		uow:        uow,
		eventRepo:  eventRepo,
		outboxRepo: outboxRepo,
		svc:        svc,
	}
}

func (h *UpdateTransactionHandler) Handle(
	ctx context.Context,
	c commands.UpdateTransactionCommand) (dto.UpdateDTO, error) {

	const op = "Services.TransactionService.UpdateTransaction"

	h.log.Infow("Attempting to update transaction...", "transactionID", c.TransactionID)

	updEvent := h.svc.ConvertCommandToUpdEvent(c)

	event, err := h.svc.CreateUpdateEvent(updEvent)
	if err != nil {
		h.log.Errorw("Failed to create update event", "transactionID", c.TransactionID)
		return dto.UpdateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	outbox := event.ToOutbox()

	ctx, err = h.uow.BeginWithTx(ctx)
	if err != nil {
		h.log.Error("Failed to begin transaction")
		return dto.UpdateDTO{}, fmt.Errorf("%s: %w", op, err)
	}
	defer h.uow.GracefulRollback(ctx, &err)

	if err = h.eventRepo.SaveEvent(ctx, event); err != nil {
		h.log.Errorw("Failed to save event to repository", "transactionID", c.TransactionID)
		return dto.UpdateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	if err = h.outboxRepo.SaveOutboxEvent(ctx, outbox); err != nil {
		h.log.Errorw("Failed to save outbox event to repository", "transactionID", c.TransactionID)
		return dto.UpdateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	if err = h.uow.Commit(ctx); err != nil {
		h.log.Errorw("Transaction commit failed", "transactionID", c.TransactionID)
		return dto.UpdateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	h.log.Infow("Account update event saved successfully", "transactionID", c.TransactionID)

	return dto.UpdateDTO{
		TransactionID: c.TransactionID,
		EventID:       event.EventID.String(),
	}, nil
}
