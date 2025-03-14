package commands

type CreateAccountCommand struct {
	ClientID string `json:"client_id" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

type UpdateAccountCommand struct {
	AccountID         string  `json:"account_id" binding:"required"`
	Amount            float64 `json:"amount" binding:"required"`
	BalanceUpdateType string  `json:"balance_update_type" binding:"required"`
	Status            string  `json:"status,omitempty"`
}
