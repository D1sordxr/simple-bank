package tests

import (
	"github.com/D1sordxr/simple-banking-system/internal/application/account/commands"
	"github.com/D1sordxr/simple-banking-system/internal/domain/account"
	loadLogger "github.com/D1sordxr/simple-banking-system/internal/infrastructure/app/logger"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/app/logger/handlers/designed"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/mocks"
	"github.com/stretchr/testify/mock"
)

func mockAccountDeps() *commands.Dependencies {

	logger := &loadLogger.Logger{Logger: designed.NewPrettySlog()}

	uow := &mocks.TestUoWManager{}

	mockAccountRepo := new(mocks.MockAccountRepository)
	mockAccountRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockAccountRepo.On("getByID", mock.Anything, mock.Anything).Return(account.Aggregate{}, nil)

	mockEventRepo := new(mocks.MockEventRepository)
	mockEventRepo.On("SaveEvent", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	mockOutboxRepo := new(mocks.MockOutboxRepository)
	mockOutboxRepo.On("SaveOutboxEvent", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	return &commands.Dependencies{
		Logger:            logger,
		UoWManager:        uow,
		EventRepository:   mockEventRepo,
		OutboxRepository:  mockOutboxRepo,
		AccountRepository: mockAccountRepo,
	}
}
