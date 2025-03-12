package handlers

type TransactionProcessor struct {
	MessageProc
}

//func NewTransactionProcessor(transactionSvc transaction.ProcessTransaction) *TransactionProcessor {
//	return &TransactionProcessor{transactionSvc: transactionSvc}
//}
//
//func (t *TransactionProcessor) Process(msg []byte) {
//	data := transaction.ProcessDTO{ByteData: msg}
//
//	err := t.transactionSvc.Handle(data)
//	if err != nil {
//		// log error
//	}
//}
