package client

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/client"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/models"
)

func ConvertAggregateToModel(client client.Aggregate) models.Client {
	return models.Client{
		ID:        client.ClientID,
		FullName:  client.FullName.String(),
		Email:     client.Email.String(),
		Status:    client.Status.String(),
		CreatedAt: client.CreatedAt,
	}
}
