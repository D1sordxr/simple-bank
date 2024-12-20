package tests

import (
	"encoding/json"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox"
	sharedVO "github.com/D1sordxr/simple-banking-system/internal/domain/shared/vo"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction"
	"github.com/D1sordxr/simple-banking-system/internal/domain/transaction/vo"
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

	txOutbox, err := outbox.NewTransactionOutbox(mockTxAggregate)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	t.Logf("\nTransactionAggregate: %v\nTransactionOutbox: %v\n", mockTxAggregate, txOutbox)

	marshalledPayload, err := json.Marshal(txOutbox.MessagePayload)
	if err != nil {
		t.Fatalf("failed to marshal message payload: %v", err)
	}
	t.Logf("Marshalled JSON result: %s", marshalledPayload)
}
