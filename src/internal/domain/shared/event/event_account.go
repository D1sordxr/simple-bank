package event

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/account"
)

// TODO: NewAccountEvent()

func NewAccountEvent(account account.Aggregate) (Event, error) {
	return Event{}, nil
}
