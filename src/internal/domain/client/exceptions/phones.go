package exceptions

import "errors"

var InvalidPhoneData = errors.New("invalid phone data: country, code, and number must be positive integers")
