package exceptions

import "errors"

var (
	MarshalFailed  = errors.New("failed to marshal data to JSON")
	InvalidPayload = errors.New("invalid JSON payload")
)
