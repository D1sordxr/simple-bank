package account

type UpdateEvents []UpdateEvent

type UpdateEvent struct {
	AccountID         string  `json:"account_id"`
	Amount            float64 `json:"amount"`
	BalanceUpdateType string  `json:"balance_update_type"`
	TransactionID     string  `json:"transaction_id"`
}
