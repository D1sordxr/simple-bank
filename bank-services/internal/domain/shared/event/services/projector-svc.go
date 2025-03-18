package services

import (
	"encoding/json"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
)

type ProjectorDomainSvc struct{}

func (ProjectorDomainSvc) ParseEvent(data []byte) (e event.Event, err error) {
	const op = "SharedServices.DomainProjectorSvc.ParseEvent"

	err = json.Unmarshal(data, &e)
	if err != nil {
		return event.Event{}, fmt.Errorf("%s: %w", op, err)
	}

	return e, nil
}
