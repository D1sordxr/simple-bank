package handlers

import "LearningArch/bank-workers/internal/application/transaction"

type TransactionProcessor struct {
	transactionSvc transaction.ProcessTransaction
}

func NewTransactionProcessor(transactionSvc transaction.ProcessTransaction) *TransactionProcessor {
	return &TransactionProcessor{transactionSvc: transactionSvc}
}

func (t *TransactionProcessor) Process(msg []byte) {
	data := transaction.ProcessDTO{ByteData: msg}

	err := t.transactionSvc.Handle(data)
	if err != nil {
		// log error
	}
}
