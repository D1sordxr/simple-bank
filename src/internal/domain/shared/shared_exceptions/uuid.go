package shared_exceptions

import "errors"

var (
	InvalidUUID      = errors.New("client ID cannot be nil")
	ClientIDNotFound = errors.New("client with ID %s not found")
)
