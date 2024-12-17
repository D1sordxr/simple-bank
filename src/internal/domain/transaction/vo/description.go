package vo

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction/exceptions"
	"strings"
)

const maxDescriptionLength = 255

type Description struct {
	Value string
}

func NewDescription(value string) (Description, error) {
	if len(value) == 0 {
		return Description{}, nil
	}
	value = strings.TrimSpace(value)
	if len(value) > maxDescriptionLength {
		return Description{}, exceptions.InvalidLength
	}
	return Description{Value: value}, nil
}

func (d Description) String() string {
	return d.Value
}

func (d Description) IsEmpty() bool {
	return d.Value == ""
}
