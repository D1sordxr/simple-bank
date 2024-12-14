package client

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/entity"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/exceptions"
	"github.com/google/uuid"
)

type Aggregate struct {
	ClientID uuid.UUID
	FullName vo.FullName
	Email    vo.Email
	Phones   entity.Phones
	Status   vo.Status
}

func NewClient(clientID uuid.UUID,
	name vo.FullName,
	email vo.Email,
	phones entity.Phones,
	status vo.Status) (Aggregate, error) {

	if clientID == uuid.Nil {
		return Aggregate{}, exceptions.InvalidUUID
	}

	return Aggregate{
		ClientID: clientID,
		FullName: name,
		Email:    email,
		Phones:   phones,
		Status:   status,
	}, nil
}
