package account

type Projection struct {
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
	Status    string  `json:"status"`
}
