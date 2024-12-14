package vo

import "errors"

const (
	StatusActive    = "active"
	StatusClosed    = "closed"
	StatusSuspended = "suspended"
)

type Status struct {
	CurrentStatus string
}

func NewStatus(status string) (Status, error) {
	return Status{CurrentStatus: StatusActive}, nil
}

func (s *Status) String() string {
	return s.CurrentStatus
}

// SetStatus allows changing the status, ensuring it is valid
func (s *Status) SetStatus(status string) error {
	if !isValidStatus(status) {
		return errors.New("invalid status")
	}
	s.CurrentStatus = status
	return nil
}

// IsActive checks if the current status is "active"
func (s *Status) IsActive() bool {
	return s.CurrentStatus == StatusActive
}

// IsClosed checks if the current status is "closed"
func (s *Status) IsClosed() bool {
	return s.CurrentStatus == StatusClosed
}

// IsSuspended checks if the current status is "suspended"
func (s *Status) IsSuspended() bool {
	return s.CurrentStatus == StatusSuspended
}

// isValidStatus checks if the given status is valid
func isValidStatus(status string) bool {
	switch status {
	case StatusActive, StatusClosed, StatusSuspended:
		return true
	default:
		return false
	}
}
