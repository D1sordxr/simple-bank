package transaction

import (
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction/vo"
	"time"
)

func NewTransactionCreatedEvent(
	id sharedVO.UUID,
	sourceAccountID, destinationAccountID *sharedVO.UUID,
	currency sharedVO.Currency,
	amount sharedVO.Money,
	status vo.TransactionStatus,
	txType vo.Type,
	description *vo.Description) transaction.TransactionCreatedEvent {

	return transaction.TransactionCreatedEvent{
		ID:                   id,
		SourceAccountID:      sourceAccountID,
		DestinationAccountID: destinationAccountID,
		Currency:             currency,
		Amount:               amount,
		Status:               status,
		Type:                 txType,
		Description:          description,
		Timestamp:            time.Now(),
	}
}
