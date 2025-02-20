package unit_of_work

import "errors"

var (
	ErrTxStartFailed = errors.New("failed to start transaction")
	ErrNoCommitTx    = errors.New("no transaction to commit")
	ErrCommitTx      = errors.New("failed to commit transaction")
	ErrNoRollbackTx  = errors.New("no transaction to rollback")
	ErrRollbackTx    = errors.New("failed to rollback tx")
	ErrExecBatch     = errors.New("error while executing batch")
	ErrClosingBatch  = errors.New("error while closing batch")
)
