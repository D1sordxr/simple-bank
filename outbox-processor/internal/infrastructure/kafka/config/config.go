package config

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	Topics  Topics   `yaml:"topics"`
}

type Topics struct {
	ClientCreatedEvent      string `yaml:"client_created_event"`
	AccountCreatedEvent     string `yaml:"account_created_event"`
	TransactionCreatedEvent string `yaml:"transaction_created_event"`
}
