package processed_messages

import (
	"context"
	"errors"
	"fmt"
	"github.com/D1sordxr/packages/postgres/executor"
	"github.com/jackc/pgx/v5"
)

type DAO struct {
	Executor *executor.Manager
}

func NewDAO(exec *executor.Manager) *DAO {
	return &DAO{Executor: exec}
}

func (d *DAO) IsProcessed(ctx context.Context, msgID string) (bool, error) {
	const op = "postgres.ProcessedMessagesDAO.IsProcessed"

	conn := d.Executor.GetExecutor(ctx)

	query := `SELECT 1 FROM processed_messages WHERE id = $1`

	var exists int
	err := conn.QueryRow(ctx, query, msgID).Scan(&exists)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("%s: %w: %w", op, ErrReadingMsg, err)
	}

	return true, nil
}

func (d *DAO) SetProcessed(ctx context.Context, msgID string) error {
	const op = "postgres.ProcessedMessagesDAO.SetProcessed"

	conn := d.Executor.GetExecutor(ctx)

	query := `INSERT INTO processed_messages (id) VALUES ($1) ON CONFLICT DO NOTHING`

	_, err := conn.Exec(ctx, query, msgID)
	if err != nil {
		return fmt.Errorf("%s: %w: %w", op, ErrInsertingMsg, err)
	}

	return nil
}
