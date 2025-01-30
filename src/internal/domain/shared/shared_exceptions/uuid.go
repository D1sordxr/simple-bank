package shared_exceptions

import "errors"

var (
	InvalidUUID = errors.New("ID cannot be nil")
)
