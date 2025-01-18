package client

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/entity"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_exceptions"
	"github.com/google/uuid"
	"time"
)

type Aggregate struct {
	ClientID  uuid.UUID
	FullName  vo.FullName
	Email     vo.Email
	Phones    entity.Phones
	Status    vo.Status
	CreatedAt time.Time
}

func NewClient(clientID uuid.UUID,
	name vo.FullName,
	email vo.Email,
	phones entity.Phones,
	status vo.Status) (Aggregate, error) {

	if clientID == uuid.Nil {
		return Aggregate{}, shared_exceptions.InvalidUUID
	}

	return Aggregate{
		ClientID:  clientID,
		FullName:  name,
		Email:     email,
		Phones:    phones,
		Status:    status,
		CreatedAt: time.Now(),
	}, nil
}
