package commands

type CreateTransactionCommand struct {
	SourceAccountID      string  `json:"source_account_id"`
	DestinationAccountID string  `json:"destination_account_id"`
	Currency             string  `json:"currency" binding:"required"`
	Amount               float64 `json:"amount" binding:"required"`
	Type                 string  `json:"type" binding:"required"`
	Description          string  `json:"description"`
}
