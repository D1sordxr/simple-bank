package vo

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/client/exceptions"
	"net/mail"
	"unicode/utf8"
)

type Email struct {
	Email string `json:"email"`
}

func NewEmail(email string) (Email, error) {
	if len(email) == 0 || utf8.RuneCountInString(email) > 255 {
		return Email{}, exceptions.InvalidEmailLength
	}

	address, err := mail.ParseAddress(email)
	if err != nil {
		return Email{}, exceptions.InvalidEmailFormat
	}
	return Email{Email: address.Address}, nil
}

func (e *Email) String() string {
	return e.Email
}
