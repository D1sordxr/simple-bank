package outbox

import (
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/vo"
	"time"
)

type Outbox struct {
	OutboxID       sharedVO.UUID // Outbox unique ID
	AggregateID    int           // TODO: vo.AggregateID (Client: 1, Account: 2, Transaction: 3)
	AggregateType  string        // TODO: vo.AggregateType (Client, Account, Transaction)
	MessageType    string        // TODO: vo.MessageType (Created, Updated, Deleted)
	MessagePayload string        // TODO: vo.MessagePayload (contains marshaled JSON)
	CreatedAt      time.Time     // Creation time
	Status         string        // TODO: vo.OutboxStatus (Pending, Processed, Failed)
	RetryCount     int           // Number of retries
}
