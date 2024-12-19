package outbox

import (
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox/vo"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction"
	"time"
)

func NewTransactionOutbox(tx transaction.Aggregate) (Outbox, error) {
	outboxID := sharedVO.NewUUID()
	aggregateID := vo.NewTransactionAggregateID()
	aggregateType := vo.NewTransactionAggregateType()
	messageType, err := vo.NewMessageType(vo.TypeCreated)
	if err != nil {
		return Outbox{}, err
	}
	messagePayload, err := vo.NewMessagePayload(tx)
	if err != nil {
		return Outbox{}, err
	}
	outboxStatus, err := vo.NewOutboxStatus(vo.StatusPending)
	if err != nil {
		return Outbox{}, err
	}
	creationTime := time.Now()
	return Outbox{
		OutboxID:       outboxID,
		AggregateID:    aggregateID,
		AggregateType:  aggregateType,
		MessageType:    messageType,
		MessagePayload: messagePayload,
		Status:         outboxStatus,
		CreatedAt:      creationTime,
	}, nil
}
