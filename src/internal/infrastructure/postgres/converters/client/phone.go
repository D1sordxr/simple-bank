package client

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/entity"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/models"
)

func ConvertPhoneEntityToModel(entity entity.Phone) models.Phone {
	return models.Phone{
		ID:          entity.PhoneID,
		ClientID:    entity.ClientID,
		PhoneNumber: entity.String(),
		Country:     entity.Country,
		Code:        entity.Code,
		Number:      entity.Number,
		CreatedAt:   entity.CreatedAt,
	}
}
