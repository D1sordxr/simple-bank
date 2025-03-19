package services

import (
	"errors"
)

var (
	ErrParsingMsg = errors.New("error parsing message")
)

var (
	ErrDifferentTxIDs    = errors.New("different transaction IDs")
	ErrStatusesEqual     = errors.New("statuses are equal")
	ErrTxHasStatusFailed = errors.New("transaction has status failed")
	FailedToParseUUID    = errors.New("failed to parse UUID")
)
