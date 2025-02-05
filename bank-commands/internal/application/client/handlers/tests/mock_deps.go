package tests

import (
	"github.com/D1sordxr/simple-banking-system/internal/application/client/commands"
	loadLogger "github.com/D1sordxr/simple-banking-system/internal/infrastructure/app/logger"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/app/logger/handlers/designed"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/mocks"
	"github.com/stretchr/testify/mock"
)

func mockClientDeps() *commands.Dependencies {

	logger := &loadLogger.Logger{Logger: designed.NewPrettySlog()}

	uow := &mocks.TestUoWManager{}

	mockClientRepo := new(mocks.MockClientRepository)
	mockClientRepo.On("Exists", mock.Anything, mock.Anything).Return(nil)
	mockClientRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	mockEventRepo := new(mocks.MockEventRepository)
	mockEventRepo.On("SaveEvent", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	mockOutboxRepo := new(mocks.MockOutboxRepository)
	mockOutboxRepo.On("SaveOutboxEvent", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	return &commands.Dependencies{
		Logger:           logger,
		UoWManager:       uow,
		EventRepository:  mockEventRepo,
		OutboxRepository: mockOutboxRepo,
		ClientRepository: mockClientRepo,
	}
}
