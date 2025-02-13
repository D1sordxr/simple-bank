package application

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/client"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/transaction"
)

type Services struct {
	ClientService      *client.Service
	AccountService     *account.Service
	TransactionService *transaction.Service
}

func NewApplicationServices(
	client *client.Service,
	account *account.Service,
	transaction *transaction.Service,
) *Services {
	return &Services{
		ClientService:      client,
		AccountService:     account,
		TransactionService: transaction,
	}
}
