package transaction

import "errors"

var (
	ErrReadingTransaction  = errors.New("error reading transaction")
	ErrTransactionNotFound = errors.New("transaction not found")
)
