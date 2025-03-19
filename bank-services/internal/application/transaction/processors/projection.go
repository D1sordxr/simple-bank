package processors

import (
	"context"
	"fmt"
	sharedInterfaces "github.com/D1sordxr/simple-bank/bank-services/internal/application/shared/interfaces"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/dto"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction/interfaces"
)

type ProjectionProcessor struct {
	sharedSvc sharedInterfaces.ProjectionDomainSvc
	svc       interfaces.ProjectionDomainSvc
	dao       interfaces.TransactionDAO
}

func NewProjectionProcessor(
	sharedSvc sharedInterfaces.ProjectionDomainSvc,
) *ProjectionProcessor {
	return &ProjectionProcessor{
		sharedSvc: sharedSvc,
	}
}

func (p *ProjectionProcessor) Process(ctx context.Context, dto dto.ProcessDTO) error {
	const op = "application.transaction.processors.ProjectionProcessor.Process"

	event, err := p.sharedSvc.ParseEvent(dto.Value)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// todo: insert into table processed_events

	upd, err := p.svc.ParseUpdateEvent([]byte(event.Payload.Payload))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// todo: dao.getById(upd.TransactionID)

	return nil
}
