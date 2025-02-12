package client

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/client"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/models"
)

func ConvertAggregateToModel(agg client.Aggregate) models.Client {
	return models.Client{
		ID:         agg.ClientID.Value,
		FirstName:  agg.FullName.FirstName,
		LastName:   agg.FullName.LastName,
		MiddleName: agg.FullName.MiddleName,
		Email:      agg.Email.String(),
		Status:     agg.Status.String(),
		CreatedAt:  agg.CreatedAt,
		UpdatedAt:  agg.UpdatedAt,
	}
}
