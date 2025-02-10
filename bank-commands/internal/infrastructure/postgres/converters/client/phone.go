package client

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/client/entity"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/models"
)

func ConvertPhoneEntityToModel(phone entity.Phone) models.Phone {
	return models.Phone{
		ID:          phone.PhoneID,
		ClientID:    phone.ClientID,
		PhoneNumber: phone.String(),
		CreatedAt:   phone.CreatedAt,
		UpdatedAt:   phone.UpdatedAt,
	}
}
