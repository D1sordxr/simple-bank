package vo

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/exceptions"
)

var validCurrencies = map[string]bool{
	"USD": true,
	"EUR": true,
	"RUB": true,
}

type Currency struct {
	Code string
}

func NewCurrency(code string) (Currency, error) {
	if !validCurrencies[code] {
		return Currency{}, exceptions.InvalidCurrency
	}
	return Currency{Code: code}, nil

}

func (c Currency) String() string {
	return c.Code
}
