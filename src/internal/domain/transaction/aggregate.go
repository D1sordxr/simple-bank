package transaction

import (
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/vo"
	"github.com/google/uuid"
	"time"
)

type Aggregate struct {
	TransactionID        uuid.UUID         // unique identifier for the transaction
	SourceAccountID      *uuid.UUID        // source account (nullable for deposits)
	DestinationAccountID *uuid.UUID        // destination account (nullable for withdrawals)
	Currency             sharedVO.Currency // transaction currency
	Amount               sharedVO.Money    // transaction amount
	TransactionStatus    string            // TODO: vo.TransactionStatus // status: initiated, completed, failed, canceled
	Type                 string            // TODO: vo.Type // type: transfer, deposit, withdrawal, reversal
	Description          string            // TODO: vo.Description if needed // optional transaction description
	FailureReason        *string           // reason for failure (nullable)
	Timestamp            time.Time         // time of transaction initiation
}
