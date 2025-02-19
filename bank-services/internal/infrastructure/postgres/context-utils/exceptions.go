package context_utils

import (
	"errors"
)

var (
	ErrNoTxInCtx    = errors.New("transaction not found in context")
	ErrNoBatchInCtx = errors.New("batch not found in context")
)
