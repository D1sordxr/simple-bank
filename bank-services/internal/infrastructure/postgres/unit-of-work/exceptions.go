package unit_of_work

import "errors"

var (
	ErrTxAlreadyStarted = errors.New("transaction already started")
	ErrTxStartFailed    = errors.New("failed to start transaction")
	ErrNoCommitTx       = errors.New("no transaction to commit")
	ErrCommitTx         = errors.New("failed to commit transaction")
	ErrNoRollbackTx     = errors.New("no transaction to rollback")
	ErrRollbackTx       = errors.New("failed to rollback tx")
)
