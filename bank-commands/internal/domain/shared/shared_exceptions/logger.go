package shared_exceptions

import (
	"fmt"
	"log/slog"
)

func LogError(err error) slog.Attr {
	return slog.Attr{
		Key:   "Error",
		Value: slog.StringValue(err.Error()),
	}
}

// LogAggregateCreationError accepts aggregate name and returns error string.
func LogAggregateCreationError(aggregateName string) string {
	return fmt.Sprintf("Failed to create %s aggregate", aggregateName)
}

// LogEntityCreationError accepts entity name and returns error string.
func LogEntityCreationError(entityName string) string {
	return fmt.Sprintf("Failed to create %s entity", entityName)
}

// LogVOCreationError accepts valueObject name and returns error string.
func LogVOCreationError(voName string) string {
	return fmt.Sprintf("Failed to create %s value object", voName)
}

// LogEventCreationError returns event creation error.
func LogEventCreationError() string {
	return fmt.Sprintf("Failed to create event")
}

// LogOutboxCreationError returns outbox creation error.
func LogOutboxCreationError() string {
	return fmt.Sprintf("Failed to create outbox")
}

// LogErrorAsString logs error as string.
func LogErrorAsString(err error) string {
	return err.Error()
}

// LogTransactionCreationError logs an error that occurred during transaction creation.
func LogTransactionCreationError(err error) string {
	return fmt.Sprintf("Error during transaction creation %v", err)
}

// LogTransactionRollbackError logs an error that occurred during transaction rollback.
func LogTransactionRollbackError(err error) string {
	return fmt.Sprintf("Error during transaction rollback %v", err)

}

// LogTransactionCommitError logs an error that occurred during transaction commit.
func LogTransactionCommitError(err error) string {
	return fmt.Sprintf("Error during transaction commit %v", err)
}
