package queries

type GetByIDAccountCommand struct {
	AccountID string `json:"account_id" binding:"required"`
}
