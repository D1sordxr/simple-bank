package client

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/client/entity"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/client/vo"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/shared_exceptions"
	sharedVO "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/shared_vo"
	"github.com/google/uuid"
	"time"
)

type Aggregate struct {
	ClientID  sharedVO.UUID
	FullName  vo.FullName
	Email     vo.Email
	Phones    entity.Phones
	Status    vo.Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(clientID sharedVO.UUID,
	name vo.FullName,
	email vo.Email,
	phones entity.Phones,
	status vo.Status) (Aggregate, error) {

	if clientID.Value == uuid.Nil {
		return Aggregate{}, shared_exceptions.InvalidUUID
	}

	return Aggregate{
		ClientID:  clientID,
		FullName:  name,
		Email:     email,
		Phones:    phones,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
