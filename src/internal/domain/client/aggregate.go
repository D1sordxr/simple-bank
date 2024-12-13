package client

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/entity"
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/vo"
	"github.com/google/uuid"
)

type Aggregate struct {
	ClientID uuid.UUID
	FullName vo.FullName
	Email    vo.Email
	Phones   entity.Phones
	Status   vo.Status
}
