package transaction

import (
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction/exceptions"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction/vo"
	"github.com/google/uuid"
	"time"
)

type Aggregate struct {
	TransactionID        uuid.UUID            // unique identifier for the transaction
	SourceAccountID      *uuid.UUID           // source account (nullable for deposits)
	DestinationAccountID *uuid.UUID           // destination account (nullable for withdrawals)
	Currency             sharedVO.Currency    // transaction currency
	Amount               sharedVO.Money       // transaction amount
	TransactionStatus    vo.TransactionStatus // status: initiated, completed, failed, canceled
	Type                 vo.Type              // type: transfer, deposit, withdrawal, reversal
	Description          *vo.Description      // optional transaction description
	FailureReason        *string              // reason for failure (nullable)
	Timestamp            time.Time            // time of transaction initiation
}

func NewTransaction(
	txID uuid.UUID,
	sourceAccountID *uuid.UUID,
	destinationAccountID *uuid.UUID,
	currency sharedVO.Currency,
	amount sharedVO.Money,
	txType vo.Type,
	description *vo.Description) (Aggregate, error) {
	if txID == uuid.Nil {
		return Aggregate{}, exceptions.InvalidTxID
	}
	if sourceAccountID == nil && txType.Value != vo.DepositType {
		return Aggregate{}, exceptions.NoSourceWithDepositType
	}
	if destinationAccountID == nil && txType.Value != vo.WithdrawalType {
		return Aggregate{}, exceptions.NoDestinationWithWithdrawalType
	}
	if sourceAccountID != nil && destinationAccountID != nil && *sourceAccountID == *destinationAccountID {
		return Aggregate{}, exceptions.EqualUUIDs
	}

	status := vo.NewTransactionStatus()
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
		Timestamp:            time.Now(),
	}, nil
}
