package vo

const (
	ClientAggregateID      = 1
	AccountAggregateID     = 2
	TransactionAggregateID = 3
)

type AggregateID struct {
	ID int
}

func NewClientAggregateID() AggregateID {
	return AggregateID{ID: ClientAggregateID}
}

func NewAccountAggregateID() AggregateID {
	return AggregateID{ID: AccountAggregateID}
}

func NewTransactionAggregateID() AggregateID {
	return AggregateID{ID: TransactionAggregateID}
}
