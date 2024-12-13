package client

import (
	"LearningArch/internal/domain/client/entity"
	"LearningArch/internal/domain/client/vo"
	"github.com/google/uuid"
)

type Aggregate struct {
	ClientID uuid.UUID
	FullName vo.FullName
	Email    vo.Email
	Phones   entity.Phones
	Status   vo.Status
}
