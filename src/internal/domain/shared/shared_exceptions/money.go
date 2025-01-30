package shared_exceptions

import "errors"

var (
	InvalidMoney = errors.New("invalid amount of money")
)
