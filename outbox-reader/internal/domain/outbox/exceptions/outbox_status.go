package exceptions

import "errors"

var InvalidOutboxStatus = errors.New("invalid outbox status")
