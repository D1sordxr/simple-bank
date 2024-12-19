package outbox

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction"
)

// TODO: NewTransactionOutbox()

func NewTransactionOutbox(transaction transaction.Aggregate) Outbox {
	return Outbox{}
}
