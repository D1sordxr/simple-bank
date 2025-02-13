package vo

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/client/exceptions"
)

const (
	StatusActive  = "active"
	StatusArchive = "archive"
)

type Status struct {
	Status string
}

func NewStatus() Status {
	return Status{Status: StatusActive}
}

func (s *Status) ChangeStatus(status string) error {
	if status != StatusActive && status != StatusArchive {
		return exceptions.InvalidStatus
	}
	s.Status = status
	return nil
}

func (s *Status) String() string {
	return s.Status
}
