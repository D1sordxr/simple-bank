package vo

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/exceptions"
	"github.com/google/uuid"
)

type UUID struct {
	Value uuid.UUID
}

// NewUUID creates a new UUID
func NewUUID() UUID {
	return UUID{Value: uuid.New()}
}

// NewUUIDFromString parses a string into a UUID or returns an error
func NewUUIDFromString(input string) (UUID, error) {
	parsed, err := uuid.Parse(input)
	if err != nil {
		return UUID{}, exceptions.InvalidUUID
	}
	return UUID{Value: parsed}, nil
}

// NewPointerUUID creates a new pointer to UUID
func NewPointerUUID() *UUID {
	u := NewUUID()
	return &u
}

// NewPointerUUIDFromString parses a string into a pointer to UUID or returns an error
func NewPointerUUIDFromString(input string) (*UUID, error) {
	u, err := NewUUIDFromString(input)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// String returns the string representation of the UUID
func (u UUID) String() string {
	return u.Value.String()
}

// IsNil checks if the UUID is nil
func (u UUID) IsNil() bool {
	return u.Value == uuid.Nil
}
