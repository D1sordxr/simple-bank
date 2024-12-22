package outbox

import (
	vo2 "github.com/D1sordxr/simple-banking-system/internal/domain/shared/event/vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox/vo"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_vo"
	"time"
)

// TODO: RetryCount	(int)	number of retries

type Outbox struct {
	OutboxID       sharedVO.UUID      // Outbox unique ID
	AggregateID    sharedVO.UUID      // References to aggregate unique ID
	AggregateType  vo2.AggregateType  // Client, Account, Transaction
	MessageType    vo2.MessageType    // Created, Updated, Deleted
	MessagePayload vo2.MessagePayload // Contains marshalled JSON
	Status         vo.OutboxStatus    // Pending, Processed, Failed
	CreatedAt      time.Time          // Creation time
}
