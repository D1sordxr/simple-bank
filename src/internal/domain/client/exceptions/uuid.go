package exceptions

import "errors"

var InvalidUUID = errors.New("client ID cannot be nil")
