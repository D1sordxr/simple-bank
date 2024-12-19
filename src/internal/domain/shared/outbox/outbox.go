package outbox

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox/vo"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/vo"
	"time"
)

// TODO: RetryCount	(int)	number of retries

type Outbox struct {
	OutboxID       sharedVO.UUID     // Outbox unique ID
	AggregateID    vo.AggregateID    // Client: 1, Account: 2, Transaction: 3
	AggregateType  vo.AggregateType  // Client, Account, Transaction
	MessageType    vo.MessageType    // Created, Updated, Deleted
	MessagePayload vo.MessagePayload // Contains marshalled JSON
	Status         vo.OutboxStatus   // Pending, Processed, Failed
	CreatedAt      time.Time         // Creation time
}
