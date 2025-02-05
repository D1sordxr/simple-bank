package vo

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction/exceptions"
)

const (
	StatusInitiated = "initiated"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
	StatusCanceled  = "canceled"
)

var validStatuses = map[string]bool{
	StatusInitiated: true,
	StatusCompleted: true,
	StatusFailed:    true,
	StatusCanceled:  true,
}

type TransactionStatus struct {
	Status string
}

func NewTransactionStatus() TransactionStatus {
	return TransactionStatus{Status: StatusInitiated}
}

// NewTransactionStatusWithValue creates a new TransactionStatus with the given value.
func NewTransactionStatusWithValue(status string) (TransactionStatus, error) {
	if !isValidStatus(status) {
		return TransactionStatus{}, exceptions.InvalidTxStatus
	}
	return TransactionStatus{Status: status}, nil
}

// IsCompleted checks if the current status is "completed".
func (ts *TransactionStatus) IsCompleted() bool {
	return ts.Status == StatusCompleted
}

// IsFailed checks if the current status is "failed".
func (ts *TransactionStatus) IsFailed() bool {
	return ts.Status == StatusFailed
}

// IsCanceled checks if the current status is "canceled".
func (ts *TransactionStatus) IsCanceled() bool {
	return ts.Status == StatusCanceled
}

// UpdateStatus safely updates the transaction status.
func (ts *TransactionStatus) UpdateStatus(newStatus string) error {
	if !isValidStatus(newStatus) {
		return exceptions.InvalidTxStatus
	}

	// Example: add rules for allowed transitions
	if ts.Status == StatusCompleted {
		return exceptions.FailedToUpdateStatus
	}
	ts.Status = newStatus
	return nil
}

// String returns the string representation of the status.
func (ts *TransactionStatus) String() string {
	return ts.Status
}

// isValidStatus checks if the given status is valid.
func isValidStatus(status string) bool {
	return validStatuses[status]
}

// AllowedStatuses returns a list of all valid transaction statuses.
func AllowedStatuses() []string {
	return []string{StatusInitiated, StatusCompleted, StatusFailed, StatusCanceled}
}
