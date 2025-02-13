package vo

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/exceptions"
)

const (
	TransferType   = "transfer"
	DepositType    = "deposit"
	WithdrawalType = "withdrawal"
	ReversalType   = "reversal"
)

var validTypes = map[string]bool{
	TransferType:   true,
	DepositType:    true,
	WithdrawalType: true,
	ReversalType:   true,
}

type Type struct {
	Value string
}

func NewType(currentType string) (Type, error) {
	if !isValidType(currentType) {
		return Type{}, exceptions.InvalidTxType
	}
	return Type{Value: currentType}, nil
}

// String returns the string representation of the type.
func (t Type) String() string {
	return t.Value
}

// IsTransfer checks if the type is "transfer".
func (t Type) IsTransfer() bool {
	return t.Value == TransferType
}

// IsDeposit checks if the type is "deposit".
func (t Type) IsDeposit() bool {
	return t.Value == DepositType
}

// IsWithdrawal checks if the type is "withdrawal".
func (t Type) IsWithdrawal() bool {
	return t.Value == WithdrawalType
}

// IsReversal checks if the type is "reversal".
func (t Type) IsReversal() bool {
	return t.Value == ReversalType
}

// isValidType checks if the given type is valid.
func isValidType(t string) bool {
	return validTypes[t]
}

// AllowedTypes returns a list of all valid types.
func AllowedTypes() []string {
	return []string{TransferType, DepositType, WithdrawalType, ReversalType}
}
