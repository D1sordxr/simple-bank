package commands

type CreateAccountCommand struct {
	ClientID string `json:"client_id" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}
