package vo

const (
	ClientAggregateType      = "Client"
	AccountAggregateType     = "Account"
	TransactionAggregateType = "Transaction"
)

type AggregateType struct {
	Type string
}

func NewClientAggregateType() AggregateType {
	return AggregateType{Type: ClientAggregateType}
}

func NewAccountAggregateType() AggregateType {
	return AggregateType{Type: AccountAggregateType}
}

func NewTransactionAggregateType() AggregateType {
	return AggregateType{Type: TransactionAggregateType}
}
