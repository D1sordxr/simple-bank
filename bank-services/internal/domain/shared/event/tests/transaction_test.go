package tests

import (
	"encoding/json"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event"
	sharedVO "github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/shared_vo"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/transaction/vo"
	"testing"
	"time"
)

func TestSuccessNewTransactionOutbox(t *testing.T) {
	currency, err := sharedVO.NewCurrency("USD")
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	depositType, err := vo.NewType(vo.DepositType)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	description, err := vo.NewDescription("testing description")
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	mockTxAggregate := transaction.Aggregate{
		TransactionID:        sharedVO.NewUUID(),
		SourceAccountID:      sharedVO.NewPointerUUID(),
		DestinationAccountID: nil,
		Currency:             currency,
		Amount:               sharedVO.Money{Value: 2000},
		TransactionStatus:    vo.NewTransactionStatus(),
		Type:                 depositType,
		Description:          description,
		FailureReason:        nil,
		Timestamp:            time.Now(),
	}

	txOutbox, err := event.NewTransactionCreatedEvent(mockTxAggregate)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	t.Logf("\nTransactionAggregate: %v\nTransactionOutbox: %v\n", mockTxAggregate, txOutbox)

	marshalledPayload, err := json.Marshal(txOutbox.Payload)
	if err != nil {
		t.Fatalf("failed to marshal message payload: %v", err)
	}
	t.Logf("Marshalled JSON result: %s", marshalledPayload)
}
