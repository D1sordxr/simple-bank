package event

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/client"
)

// TODO: NewClientEvent()

func NewClientEvent(client client.Aggregate) (Event, error) {
	return Event{}, nil
}
