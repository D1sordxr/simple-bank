package unit_of_work

import (
	"context"
	"errors"
	"fmt"
	"github.com/D1sordxr/simple-bank/bank-services/internal/infrastructure/app/logger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const UoWContextKey = "uow"

// UnitOfWork определяет интерфейс для работы с транзакциями и батчами.
type UnitOfWork interface {
	BeginTxWithBatch(ctx context.Context) (context.Context, error)
	BeginTx(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type UnitOfWorkImpl struct {
	Logger *logger.Logger
	Pool   *pgxpool.Pool
	Batch  *pgx.Batch
	Tx     pgx.Tx
}

func NewUnitOfWork(
	logger *logger.Logger,
	pool *pgxpool.Pool,
	batch *pgx.Batch,
	tx pgx.Tx,
) *UnitOfWorkImpl {
	return &UnitOfWorkImpl{
		Logger: logger,
		Pool:   pool,
		Batch:  batch,
		Tx:     tx,
	}
}

var (
	ErrTxAlreadyStarted = errors.New("transaction already started")
	ErrTxStartFailed    = errors.New("failed to start transaction")
	ErrSendingBatch     = errors.New("failed to close batch")
	ErrCommitTx         = errors.New("failed to commit transaction")
	ErrNoRollbackTx     = errors.New("no transaction to rollback")
	ErrRollbackTx       = errors.New("failed to rollback tx")
	ErrUoWNotFound      = errors.New("unit of work not found in context")
	ErrInvalidUoWType   = errors.New("invalid unit of work type in context")
)

func (u *UnitOfWorkImpl) BeginTxWithBatch(ctx context.Context) (context.Context, error) {
	const op = "postgres.UnitOfWork.BeginTxWithBatch"
	log := u.Logger.With(u.Logger.String("operation", op))

	if u.Tx != nil {
		return ctx, fmt.Errorf("%s: %w", op, ErrTxAlreadyStarted)
	}

	log.Info("Starting new transaction with batch")

	tx, err := u.Pool.Begin(ctx)
	if err != nil {
		log.Error("Failed to start transaction", "error", err)
		return context.Background(), fmt.Errorf("%s: %w: %v", op, ErrTxStartFailed, err)
	}
	u.Tx = tx

	batch := &pgx.Batch{}
	u.Batch = batch

	ctx = context.WithValue(ctx, UoWContextKey, u)

	return ctx, nil
}

func (u *UnitOfWorkImpl) BeginTx(ctx context.Context) (context.Context, error) {
	const op = "postgres.UnitOfWork.BeginTx"
	log := u.Logger.With(u.Logger.String("operation", op))

	if u.Tx != nil {
		return ctx, fmt.Errorf("%s: %w", op, ErrTxAlreadyStarted)
	}

	log.Info("Starting new transaction")

	tx, err := u.Pool.Begin(ctx)
	if err != nil {
		log.Error("Failed to start transaction", "error", err)
		return context.Background(), fmt.Errorf("%s: %w: %v", op, ErrTxStartFailed, err)
	}
	u.Tx = tx

	ctx = context.WithValue(ctx, UoWContextKey, u)

	return ctx, nil
}

func (u *UnitOfWorkImpl) Commit(ctx context.Context) error {
	const op = "postgres.UnitOfWork.Commit"
	log := u.Logger.With(u.Logger.String("operation", op))

	uow, err := parseContext(ctx)
	if err != nil {
		log.Error("Context parsing failed", "error", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	if uow.Batch != nil {
		br := uow.Tx.SendBatch(ctx, uow.Batch)
		if err = br.Close(); err != nil {
			log.Error("Failed to close batch", "error", err)
			return fmt.Errorf("%s: %w: %v", op, ErrSendingBatch, err)
		}
	}

	if err = uow.Tx.Commit(ctx); err != nil {
		log.Error("Failed to commit transaction", "error", err)
		return fmt.Errorf("%s: %w: %v", op, ErrCommitTx, err)
	}

	uow.Tx = nil
	uow.Batch = nil

	log.Info("Transaction committed successfully")
	return nil
}

func (u *UnitOfWorkImpl) Rollback(ctx context.Context) error {
	const op = "postgres.UnitOfWork.Rollback"
	log := u.Logger.With(u.Logger.String("operation", op))

	if u.Tx == nil {
		log.Error("No transaction to rollback")
		return fmt.Errorf("%s: %w", op, ErrNoRollbackTx)
	}

	if err := u.Tx.Rollback(ctx); err != nil {
		log.Error("Failed to rollback transaction", "error", err)
		return fmt.Errorf("%s: %w: %v", op, ErrRollbackTx, err)
	}

	u.Tx = nil
	u.Batch = nil

	log.Info("Transaction rolled back")
	return nil
}

func parseContext(ctx context.Context) (*UnitOfWorkImpl, error) {
	uow := ctx.Value(UoWContextKey)
	if uow == nil {
		return nil, ErrUoWNotFound
	}

	uowImpl, ok := uow.(*UnitOfWorkImpl)
	if !ok {
		return nil, ErrInvalidUoWType
	}

	return uowImpl, nil
}
