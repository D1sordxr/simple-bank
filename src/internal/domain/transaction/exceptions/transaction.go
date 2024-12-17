package exceptions

import (
	"errors"
)

var (
	InvalidTxID                     = errors.New("transaction ID cannot be nil")
	NoSourceWithDepositType         = errors.New("source account is required for non-deposit transactions")
	NoDestinationWithWithdrawalType = errors.New("destination account is required for non-withdrawal transactions")
	EqualUUIDs                      = errors.New("source and destination accounts cannot be the same")
)
