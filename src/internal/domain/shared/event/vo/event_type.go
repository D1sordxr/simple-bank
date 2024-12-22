package vo

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event/exceptions"
)

const (
	TypeCreated = "Created"
	TypeUpdated = "Updated"
	TypeDeleted = "Deleted"
)

type EventType struct {
	Type string
}

func NewEventType(t string) (EventType, error) {
	if !isValidMessageType(t) {
		return EventType{}, exceptions.InvalidEventType
	}
	return EventType{Type: t}, nil
}

func isValidMessageType(t string) bool {
	return t == TypeCreated || t == TypeUpdated || t == TypeDeleted
}

func (e EventType) String() string {
	return e.Type
}
