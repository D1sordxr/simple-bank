package queries

type OutboxQuery struct {
	AggregateType string
	Status        string
	Limit         int
}
