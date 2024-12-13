package vo

import (
	"fmt"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/exceptions"
	"regexp"
)

var NameRegex = regexp.MustCompile(`^[A-Z][a-z]{1,254}$`)

type FullName struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func NewFullName(firstName, middleName, lastName string) (FullName, error) {
	if !NameRegex.MatchString(firstName) {
		return FullName{}, exceptions.InvalidFullName
	}
	if !NameRegex.MatchString(middleName) {
		return FullName{}, exceptions.InvalidFullName
	}
	if !NameRegex.MatchString(lastName) {
		return FullName{}, exceptions.InvalidFullName
	}

	return FullName{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
	}, nil
}

func (f *FullName) String() string {
	return fmt.Sprintf("%s %s %s", f.FirstName, f.LastName, f.MiddleName)
}
