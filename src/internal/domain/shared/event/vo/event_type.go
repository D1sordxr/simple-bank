package vo

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox/exceptions"
)

const (
	TypeCreated = "Created"
	TypeUpdated = "Updated"
	TypeDeleted = "Deleted"
)

type MessageType struct {
	Type string
}

func NewEventType(t string) (MessageType, error) {
	if !isValidMessageType(t) {
		return MessageType{}, exceptions.InvalidMessageType
	}
	return MessageType{Type: t}, nil
}

func isValidMessageType(t string) bool {
	return t == TypeCreated || t == TypeUpdated || t == TypeDeleted
}

func (m MessageType) String() string {
	return m.Type
}
