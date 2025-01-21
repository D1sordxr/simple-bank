package shared_exceptions

import (
	"fmt"
	"log/slog"
)

func LogError(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

// LogEntityCreationError accepts entity name and returns error string
func LogEntityCreationError(entityName string) string {
	return fmt.Sprintf("Failed to create %s value object", entityName)
}

// LogVOCreationError accepts entity name and returns error string
func LogVOCreationError(voName string) string {
	return fmt.Sprintf("Failed to create %s value object", voName)
}
