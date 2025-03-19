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

	// todo: select table processed_events where event_id = event.ID
	// todo: if exists, return nil

	upd, err := p.svc.ParseUpdateEvent([]byte(event.Payload.Payload))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	oldModel, err := p.dao.GetProjection(ctx, upd.TransactionID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	oldProjection := p.svc.ConvertModelToProjection(oldModel)

	newProjection, err := p.svc.UpdateProjection(oldProjection, upd)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	newModel, err := p.svc.ConvertProjectionToModel(newProjection)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = p.dao.UpdateProjection(ctx, newModel)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// todo: insert into processed_events (event_id) values (event.ID)

	return nil
}
