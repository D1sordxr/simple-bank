package tests

import (
	"github.com/D1sordxr/simple-banking-system/internal/application/transaction/dependencies"
	loadLogger "github.com/D1sordxr/simple-banking-system/internal/infrastructure/app/logger"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/app/logger/handlers/designed"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/mocks"
	"github.com/stretchr/testify/mock"
)

func mockClientDeps() *dependencies.Dependencies {

	logger := &loadLogger.Logger{Logger: designed.NewPrettySlog()}

	uow := &mocks.TestUoWManager{}

	mockTxRepo := new(mocks.MockTransactionRepository)
	mockTxRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	mockEventRepo := new(mocks.MockEventRepository)
	mockEventRepo.On("SaveEvent", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	mockOutboxRepo := new(mocks.MockOutboxRepository)
	mockOutboxRepo.On("SaveOutboxEvent", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	return &dependencies.Dependencies{
		Logger:                logger,
		UoWManager:            uow,
		EventRepository:       mockEventRepo,
		OutboxRepository:      mockOutboxRepo,
		TransactionRepository: mockTxRepo,
	}
}
