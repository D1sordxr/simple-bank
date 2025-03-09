package commands

type CreateAccountCommand struct {
	ClientID string `json:"client_id" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

type UpdateAccountCommand struct {
	AccountID         string
	Amount            float64
	BalanceUpdateType string `json:"balance_update_type" binding:"required"`
}
