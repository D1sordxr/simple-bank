package transaction

type Projection struct {
	TransactionID string  `json:"transaction_id"`
	Status        string  `json:"transaction_status"`
	FailureReason *string `json:"failure_reason,omitempty"`
}
