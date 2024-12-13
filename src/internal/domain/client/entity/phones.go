package entity

import (
	"fmt"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/exceptions"
)

type Phone struct {
	Country int
	Code    int
	Number  int
}

type Phones []Phone

func NewPhones(phoneData []map[string]int) (Phones, error) {
	phones := make(Phones, 0, len(phoneData))

	for _, data := range phoneData {
		if data["country"] <= 0 || data["code"] <= 0 || data["number"] <= 0 {
			return nil, exceptions.InvalidPhoneData
		}

		phone := Phone{
			Country: data["country"],
			Code:    data["code"],
			Number:  data["number"],
		}

		phones = append(phones, phone)
	}

	return phones, nil
}

func (p Phones) Read() []string {
	phones := make([]string, 0, len(p))

	for _, data := range p {
		phone := fmt.Sprintf("+%v(%v)%v", data.Country, data.Code, data.Number)
		phones = append(phones, phone)
	}

	return phones
}
