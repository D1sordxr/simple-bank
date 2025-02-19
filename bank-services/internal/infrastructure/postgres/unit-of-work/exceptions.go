package unit_of_work

import "errors"

var (
	ErrTxAlreadyStarted = errors.New("transaction already started")
	ErrTxStartFailed    = errors.New("failed to start transaction")
	ErrSendingBatch     = errors.New("failed to close batch")
	ErrCommitTx         = errors.New("failed to commit transaction")
	ErrNoRollbackTx     = errors.New("no transaction to rollback")
	ErrRollbackTx       = errors.New("failed to rollback tx")
	ErrUoWNotFound      = errors.New("unit of work not found in context")
	ErrInvalidUoWType   = errors.New("invalid unit of work type in context")
)
