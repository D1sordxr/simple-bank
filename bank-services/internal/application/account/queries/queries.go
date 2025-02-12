package queries

type GetByIDAccountQuery struct {
	AccountID string `json:"account_id" binding:"required"`
}
