package unit_of_work

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/logger"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/postgres/executor"
	"github.com/jackc/pgx/v5"
)

type UnitOfWorkImpl struct {
	Logger   *logger.Logger
	Executor *executor.Executor
}

func NewUnitOfWork(
	logger *logger.Logger,
	executor *executor.Executor,
) *UnitOfWorkImpl {
	return &UnitOfWorkImpl{
		Logger:   logger,
		Executor: executor,
	}
}

// BeginWithTxBatch TODO: finish the implementation of the method
func (u *UnitOfWorkImpl) BeginWithTxBatch(ctx context.Context) (context.Context, error) {
	const op = "postgres.UnitOfWork.BeginTxWithBatch"
	log := u.Logger.With(u.Logger.String("operation", op))

	log.Info("Starting new transaction with batch")

	tx, err := u.Executor.Begin(ctx)
	if err != nil {
		log.Error("Failed to start transaction", "error", err)
		return ctx, fmt.Errorf("%s: %w: %v", op, ErrTxStartFailed, err)
	}

	batch := &pgx.Batch{}
	_ = batch

	// TODO: implement the logic for injecting the batch into the transaction
	ctx = u.Executor.InjectTx(ctx, tx)

	return ctx, nil
}

func (u *UnitOfWorkImpl) BeginWithTx(ctx context.Context) (context.Context, error) {
	const op = "postgres.UnitOfWork.BeginTx"
	log := u.Logger.With(u.Logger.String("operation", op))

	log.Info("Starting new transaction")

	tx, err := u.Executor.Begin(ctx)
	if err != nil {
		log.Error("Failed to start transaction", "error", err)
		return ctx, fmt.Errorf("%s: %w: %v", op, ErrTxStartFailed, err)
	}

	ctx = u.Executor.InjectTx(ctx, tx)

	return ctx, nil
}

func (u *UnitOfWorkImpl) Commit(ctx context.Context) error {
	const op = "postgres.UnitOfWork.Commit"
	log := u.Logger.With(u.Logger.String("operation", op))

	tx, ok := u.Executor.ExtractTx(ctx)
	if !ok {
		log.Error("No transaction to commit")
		return fmt.Errorf("%s: %w", op, ErrNoCommitTx)
	}

	if err := tx.Commit(ctx); err != nil {
		log.Error("Failed to commit transaction", "error", err)
		return fmt.Errorf("%s: %w: %v", op, ErrCommitTx, err)
	}

	log.Info("Transaction committed successfully")
	return nil
}

func (u *UnitOfWorkImpl) Rollback(ctx context.Context) error {
	const op = "postgres.UnitOfWork.Rollback"
	log := u.Logger.With(u.Logger.String("operation", op))

	tx, ok := u.Executor.ExtractTx(ctx)
	if !ok {
		log.Error("No transaction to rollback")
		return fmt.Errorf("%s: %w", op, ErrNoRollbackTx)
	}

	if err := tx.Rollback(ctx); err != nil {
		log.Error("Failed to rollback transaction", "error", err)
		return fmt.Errorf("%s: %w: %v", op, ErrRollbackTx, err)
	}

	log.Info("Transaction rolled back")
	return nil
}
