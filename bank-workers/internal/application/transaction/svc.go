package transaction

type ProcessTransactionHandler struct {
	//
}

func NewProcessTransactionHandler() *ProcessTransactionHandler {
	return &ProcessTransactionHandler{}
}

func (h *ProcessTransactionHandler) Handle(dto ProcessDTO) error {
	const op = "transaction.ProcessTransactionHandler.Handle"

	return nil
}
