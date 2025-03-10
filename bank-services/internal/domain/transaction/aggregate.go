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

// DTO TODO: use for db saving and kafka producing/consuming
type DTO struct {
	TransactionID        string  `json:"transaction_id"`
	SourceAccountID      *string `json:"source_account_id,omitempty"`
	DestinationAccountID *string `json:"destination_account_id,omitempty"`
	CurrencyCode         string  `json:"currency"`
	Amount               float64 `json:"amount"`
	Status               string  `json:"transaction_status"`
	Type                 string  `json:"type"`
	Description          string  `json:"description,omitempty"`
	FailureReason        *string `json:"failure_reason,omitempty"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
}
