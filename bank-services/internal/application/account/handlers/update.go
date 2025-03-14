package handlers

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/commands"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/dto"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/interfaces"
	sharedInterfaces "github.com/D1sordxr/simple-bank/bank-services/internal/application/shared/interfaces"
	eventRepo "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	outboxRepo "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/outbox"
)

type UpdateAccountHandler struct {
	log        sharedInterfaces.Logger
	uow        sharedInterfaces.IUnitOfWork
	eventRepo  sharedInterfaces.EventRepo
	outboxRepo sharedInterfaces.OutboxRepo
	svc        interfaces.UpdateDomainSvc
}

func NewUpdateAccountHandler(
	log sharedInterfaces.Logger,
	uow sharedInterfaces.IUnitOfWork,
	eventRepo eventRepo.Repository,
	outboxRepo outboxRepo.Repository,
	svc interfaces.UpdateDomainSvc,
) *UpdateAccountHandler {
	return &UpdateAccountHandler{
		log:        log,
		uow:        uow,
		eventRepo:  eventRepo,
		outboxRepo: outboxRepo,
		svc:        svc,
	}
}

func (h *UpdateAccountHandler) Handle(
	ctx context.Context,
	c commands.UpdateAccountCommand,
) (dto.UpdateDTO, error) {
	const op = "Services.AccountService.UpdateAccount"

	h.log.Infow("Attempting to update account...", "accountID", c.AccountID)

	event, err := h.svc.CreateUpdateEvent(c)
	if err != nil {
		h.log.Errorw("Failed to create update event", "accountID", c.AccountID)
		return dto.UpdateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	outbox := event.OutboxFromEvent()

	ctx, err = h.uow.BeginWithTx(ctx)
	if err != nil {
		h.log.Error("Failed to begin transaction")
		return dto.UpdateDTO{}, fmt.Errorf("%s: %w", op, err)
	}
	defer h.uow.GracefulRollback(ctx, &err)

	if err = h.eventRepo.SaveEvent(ctx, event); err != nil {
		h.log.Errorw("Failed to save event to repository", "accountID", c.AccountID)
		return dto.UpdateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	if err = h.outboxRepo.SaveOutboxEvent(ctx, outbox); err != nil {
		h.log.Errorw("Failed to save outbox event to repository", "accountID", c.AccountID)
		return dto.UpdateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	if err = h.uow.Commit(ctx); err != nil {
		h.log.Errorw("Transaction commit failed", "accountID", c.AccountID)
		return dto.UpdateDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	h.log.Infow("Account update event saved successfully", "accountID", c.AccountID)

	return dto.UpdateDTO{
		AccountID: c.AccountID,
		EventID:   event.EventID.String(),
	}, nil
}
