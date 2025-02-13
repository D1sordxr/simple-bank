package client

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/client/entity"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/models"
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
