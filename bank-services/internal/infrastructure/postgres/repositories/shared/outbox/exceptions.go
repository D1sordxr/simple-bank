package outbox

import "errors"

var (
	ErrFailedOutboxCreation = errors.New("failed to create outbox")
)
