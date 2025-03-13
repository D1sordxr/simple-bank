package kafka

type ConsumerTopics struct {
	ClientCreatedEvent      string `yaml:"client_created_event"`
	AccountCreatedEvent     string `yaml:"account_created_event"`
	TransactionCreatedEvent string `yaml:"transaction_created_event"`
}

type ProducerTopics struct {
	AccountBalanceUpdate string `yaml:"account_balance_update"`
}
