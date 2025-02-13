package dao

import (
	"errors"
)

var (
	InvalidTxType     = errors.New("invalid transaction type")
	QueryErr          = errors.New("query error")
	RowScanningErr    = errors.New("failed to scan row")
	StatusUpdateError = errors.New("failed to update outbox status")
)
