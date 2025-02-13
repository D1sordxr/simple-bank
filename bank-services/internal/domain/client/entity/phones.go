package entity

import (
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/client/exceptions"
	"github.com/google/uuid"
	"time"
)

type Phone struct {
	PhoneID   uuid.UUID
	ClientID  uuid.UUID
	Country   int
	Code      int
	Number    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Phones []Phone

func NewPhones(phoneData []map[string]int, clientID uuid.UUID) (Phones, error) {
	phones := make(Phones, 0, len(phoneData))

	for _, data := range phoneData {
		if data["country"] <= 0 || data["code"] <= 0 || data["number"] <= 0 {
			return nil, exceptions.InvalidPhoneData
		}

		phone := NewPhone(data, clientID)

		phones = append(phones, phone)
	}

	return phones, nil
}

func NewPhone(data map[string]int, clientID uuid.UUID) Phone {
	phoneID := uuid.New()

	phone := Phone{
		PhoneID:   phoneID,
		ClientID:  clientID,
		Country:   data["country"],
		Code:      data["code"],
		Number:    data["number"],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return phone
}

func (p Phones) Read() []string {
	phones := make([]string, 0, len(p))

	for _, data := range p {
		phone := fmt.Sprintf("+%v(%v)%v", data.Country, data.Code, data.Number)
		phones = append(phones, phone)
	}

	return phones
}

func (p *Phone) String() string {
	return fmt.Sprintf("+%v(%v)%v", p.Country, p.Code, p.Number)
}
