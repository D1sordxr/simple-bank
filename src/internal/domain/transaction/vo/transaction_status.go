package vo

// TODO: vo.TransactionStatus // status: initiated, completed, failed, canceled

const (
	statusInitiated = "initiated"
	statusCompleted = "completed"
	statusFailed    = "failed"
	statusCanceled  = "canceled"
)

type TransactionStatus struct {
	Status string
}

func NewTransactionStatus() TransactionStatus {
	return TransactionStatus{Status: statusInitiated}
}
