package client

import "errors"

var (
	ErrClientAlreadyExists  = errors.New("client already exists")
	ErrFailedToCreateClient = errors.New("failed to create client")
	ErrFailedToCreatePhone  = errors.New("failed to create client's phone")
)
