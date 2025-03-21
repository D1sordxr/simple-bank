package services

import "errors"

var (
	FailedToParseUUID            = errors.New("failed to parse uuid")
	ErrDifferentAccountIDs       = errors.New("different account ids")
	ErrAccountHasStatusClosed    = errors.New("account has status closed")
	ErrAccountHasStatusSuspended = errors.New("account has status suspended")
	ErrNotEnoughMoney            = errors.New("account does not have enough money")
	ErrNegativeBalance           = errors.New("account has negative balance")
)
