package transactionManager

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type txKey struct{}

type transactionRepo struct {
	db *pgxpool.Pool
}

func NewTransactionRepo(db *pgxpool.Pool) *transactionRepo {
	return &transactionRepo{db: db}
}

func (r *transactionRepo) PerformTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	// 1. Начинаем транзакцию в БД
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("unable to begin transaction: %v", err)
	}

	// 4. Откатываем транзакцию, если возникла ошибка
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	// Создание нового контекста с транзакцией
	ctxWithTx := injectTx(ctx, tx)

	// 2. Выполняем код транзакции
	if err = fn(ctxWithTx); err != nil {
		return err
	}

	// 3. Коммитим транзакцию в БД
	if err = tx.Commit(ctxWithTx); err != nil {
		return fmt.Errorf("unable to commit transaction: %v", err)
	}

	return nil
}

type Querier interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

// injectTx внедряет контекст транзакции в общий родительский контекст
func injectTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx вытаскивает контекст транзакции из родительского контекста
func extractTx(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(txKey{}).(pgx.Tx)
	return tx, ok
}
