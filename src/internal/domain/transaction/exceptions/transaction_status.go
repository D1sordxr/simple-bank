package exceptions

import "errors"

var (
	InvalidTxStatus      = errors.New("invalid transaction status")
	FailedToUpdateStatus = errors.New("cannot update status from completed")
)
