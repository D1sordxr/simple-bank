package transaction

import (
	sharedVO "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/shared_vo"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/exceptions"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/vo"
	"time"
)

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

func NewTransaction(
	txID sharedVO.UUID,
	sourceAccountID *sharedVO.UUID,
	destinationAccountID *sharedVO.UUID,
	currency sharedVO.Currency,
	amount sharedVO.Money,
	status vo.TransactionStatus,
	txType vo.Type,
	description *vo.Description) (Aggregate, error) {
	if txID.IsNil() {
		return Aggregate{}, exceptions.InvalidTxID
	}
	if (sourceAccountID == nil || sourceAccountID.IsNil()) && txType.Value != vo.DepositType {
		return Aggregate{}, exceptions.NoSourceWithDepositType
	}
	if (destinationAccountID == nil || destinationAccountID.IsNil()) && txType.Value != vo.WithdrawalType {
		return Aggregate{}, exceptions.NoDestinationWithWithdrawalType
	}
	// Prevent transactions between the same accounts
	if sourceAccountID != nil && destinationAccountID != nil &&
		!sourceAccountID.IsNil() && !destinationAccountID.IsNil() {
		if sourceAccountID.Value == destinationAccountID.Value {
			return Aggregate{}, exceptions.EqualUUIDs
		}
	}

	return Aggregate{
		TransactionID:        txID,
		SourceAccountID:      sourceAccountID,
		DestinationAccountID: destinationAccountID,
		Currency:             currency,
		Amount:               amount,
		TransactionStatus:    status,
		Type:                 txType,
		Description:          description,
		FailureReason:        nil,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}, nil
}
