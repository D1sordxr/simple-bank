package transaction

type UpdateEvent struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	FailureReason string `json:"failure_reason,omitempty"`
}

type UpdateEvents []UpdateEvent
