package transaction

import "time"

type Aggregate struct {
	TransactionID        sharedVO.UUID        // unique identifier for the transaction
	SourceAccountID      *sharedVO.UUID       // source account (nullable for deposits)
	DestinationAccountID *sharedVO.UUID       // destination account (nullable for withdrawals)
	Currency             sharedVO.Currency    // transaction currency
	Amount               sharedVO.Money       // transaction amount
	TransactionStatus    vo.TransactionStatus // status: initiated, completed, failed, canceled
	Type                 vo.Type              // type: transfer, deposit, withdrawal, reversal
	Description          *vo.Description      // optional transaction description
	FailureReason        *string              // reason for failure (nullable)
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
