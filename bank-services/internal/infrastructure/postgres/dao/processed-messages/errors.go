package processed_messages

import (
	"errors"
)

var (
	ErrReadingMsg   = errors.New("failed to read message")
	ErrInsertingMsg = errors.New("failed to add message")
)
