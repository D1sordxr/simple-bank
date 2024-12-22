package event

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/event/vo"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_vo"
	"time"
)

type Event struct {
	EventID       sharedVO.UUID    // Event unique ID
	AggregateID   sharedVO.UUID    // References to aggregate unique ID
	AggregateType vo.AggregateType // Client, Account, Transaction
	EventType     vo.EventType     // Created, Updated, Deleted
	Payload       vo.EventPayload  // Contains marshalled JSON
	CreatedAt     time.Time        // Creation time
}
