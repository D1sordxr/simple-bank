package exceptions

import "errors"

var (
	InvalidEmailLength = errors.New("invalid email length")
	InvalidEmailFormat = errors.New("invalid email format")
)
