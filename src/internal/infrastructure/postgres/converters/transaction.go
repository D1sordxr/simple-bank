package converters

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/models"
)

func ConvertAggregateToModel(transaction transaction.Aggregate) models.TransactionModel {
	return models.TransactionModel{
		ID:                   transaction.TransactionID.Value,
		SourceAccountID:      &transaction.SourceAccountID.Value,
		DestinationAccountID: &transaction.DestinationAccountID.Value,
		Currency:             transaction.Currency.String(),
		Amount:               transaction.Amount.Value,
		Status:               transaction.TransactionStatus.String(),
		Type:                 transaction.Type.String(),
		Description:          &transaction.Description.Value,
		CreatedAt:            transaction.Timestamp,
	}
}
