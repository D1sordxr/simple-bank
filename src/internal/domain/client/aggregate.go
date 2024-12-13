package client

import (
	"LearningArch/internal/domain/client/vo"
	"github.com/google/uuid"
)

type Aggregate struct {
	ClientID uuid.UUID
	FullName vo.FullName
	Email    vo.Email
	Phones   entities.Phones
	Status   vo.Status
}
