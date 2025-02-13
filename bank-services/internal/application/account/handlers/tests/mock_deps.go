package tests

import (
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/account/dependencies"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/account"
	loadLogger "github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/logger"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/logger/handlers/designed"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/mocks"
	"github.com/stretchr/testify/mock"
)

func mockAccountDeps() *dependencies.Dependencies {

	logger := &loadLogger.Logger{Logger: designed.NewPrettySlog()}

	uow := &mocks.TestUoWManager{}

	mockAccountRepo := new(mocks.MockAccountRepository)
	mockAccountRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockAccountRepo.On("getByID", mock.Anything, mock.Anything).Return(account.Aggregate{}, nil)

	mockEventRepo := new(mocks.MockEventRepository)
	mockEventRepo.On("SaveEvent", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	mockOutboxRepo := new(mocks.MockOutboxRepository)
	mockOutboxRepo.On("SaveOutboxEvent", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	return &dependencies.Dependencies{
		Logger:            logger,
		UoWManager:        uow,
		EventRepository:   mockEventRepo,
		OutboxRepository:  mockOutboxRepo,
		AccountRepository: mockAccountRepo,
	}
}
