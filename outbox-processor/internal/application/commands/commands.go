package commands

type OutboxCommand struct {
	ID     string
	Status string
	Topic  string
}
