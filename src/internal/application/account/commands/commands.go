package commands

type CreateAccountCommand struct {
	ClientID string `json:"client_id" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

type GetByIDAccountCommand struct {
	AccountID string `json:"account_id" binding:"required"`
}
