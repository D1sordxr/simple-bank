package unit_of_work

import (
	"context"
	"fmt"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app/logger"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres/executor"
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

func (u *UnitOfWorkImpl) BeginWithTxAndBatch(ctx context.Context) (context.Context, error) {
	const op = "postgres.UnitOfWork.BeginWithTxAndBatch"
	log := u.Logger.With(u.Logger.String("operation", op))

	log.Info("Starting new transaction with batch")

	tx, err := u.Executor.Begin(ctx)
	if err != nil {
		log.Error("Failed to start transaction", "error", err.Error())
		return ctx, fmt.Errorf("%s: %w: %v", op, ErrTxStartFailed, err)
	}

	batch := u.Executor.NewBatch()

	ctx = u.Executor.InjectTx(ctx, tx)
	ctx = u.Executor.InjectBatch(ctx, batch)

	return ctx, nil
}

func (u *UnitOfWorkImpl) BeginWithTx(ctx context.Context) (context.Context, error) {
	const op = "postgres.UnitOfWork.BeginWithTx"
	log := u.Logger.With(u.Logger.String("operation", op))

	log.Info("Starting new transaction")

	tx, err := u.Executor.Begin(ctx)
	if err != nil {
		log.Error("Failed to start transaction", "error", err.Error())
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

	if batchExecutor, ok := u.Executor.ExtractBatch(ctx); ok {
		results := tx.SendBatch(ctx, batchExecutor.Batch)
		for i := 0; i < batchExecutor.Batch.Len(); i++ {
			_, err := results.Exec()
			if err != nil {
				log.Error("Batch execution failed", "error", err.Error())
				return fmt.Errorf("%s: %w: %v", op, ErrExecBatch, err)
			}
		}

		if err := results.Close(); err != nil {
			log.Error("Failed to close batch results", "error", err.Error())
			return fmt.Errorf("%s: %w: %v", op, ErrClosingBatch, err)
		}

		log.Info("Batch executed successfully")
	}

	if err := tx.Commit(ctx); err != nil {
		log.Error("Failed to commit transaction", "error", err.Error())
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
		log.Error("Failed to rollback transaction", "error", err.Error())
		return fmt.Errorf("%s: %w: %v", op, ErrRollbackTx, err)
	}

	log.Info("Transaction rolled back")
	return nil
}

func (u *UnitOfWorkImpl) GracefulRollback(ctx context.Context, err *error) {
	const op = "postgres.UnitOfWork.GracefulRollback"
	log := u.Logger.With(u.Logger.String("operation", op))

	if r := recover(); r != nil {
		_ = u.Rollback(ctx)
		log.Error("Panic occurred, transaction rolled back", "panic", r)
		panic(r)
	}

	if err != nil && *err != nil {
		log.Error("Error occurred, transaction rolled back", "error", *err)
		_ = u.Rollback(ctx)
	}
}
