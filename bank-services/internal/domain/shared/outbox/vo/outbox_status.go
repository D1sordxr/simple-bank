package vo

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/outbox/exceptions"
)

const (
	StatusPending   = "Pending"
	StatusProcessed = "Processed"
	StatusFailed    = "Failed"
)

type OutboxStatus struct {
	Status string
}

func NewOutboxStatus(status string) (OutboxStatus, error) {
	if !isValidOutboxStatus(status) {
		return OutboxStatus{}, exceptions.InvalidOutboxStatus
	}
	return OutboxStatus{Status: status}, nil
}

func (o OutboxStatus) IsValid() bool {
	return isValidOutboxStatus(o.Status)
}

func (o OutboxStatus) String() string {
	return o.Status
}

func isValidOutboxStatus(status string) bool {
	switch status {
	case StatusPending, StatusProcessed, StatusFailed:
		return true
	default:
		return false
	}
}
