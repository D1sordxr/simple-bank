package account

import "errors"

var (
	ErrReadingAccount  = errors.New("failed to read account")
	ErrUpdatingAccount = errors.New("failed to update account")
)
