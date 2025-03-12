package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/commands"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/queries"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres/converters"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres/executor"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres/models"
	"github.com/jackc/pgx/v5"
)

type OutboxDAO struct {
	Executor *executor.Executor
}

func NewOutboxDAO(executor *executor.Executor) *OutboxDAO {
	return &OutboxDAO{Executor: executor}
}

func (dao *OutboxDAO) FetchMessages(ctx context.Context, q queries.OutboxQuery) (queries.OutboxDTOs, error) {
	const op = "postgres.OutboxDAO.FetchMessages"

	conn := dao.Executor.GetExecutor(ctx)

	query := `SELECT id, aggregate_id, aggregate_type, message_type, message_payload, status, created_at
          FROM outbox
          WHERE status = $1 AND aggregate_type = $2
          ORDER BY created_at
          LIMIT $3`

	rows, err := conn.Query(ctx, query, q.Status, q.AggregateType, q.Limit)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return queries.OutboxDTOs{}, nil
		}
		return nil, fmt.Errorf("%s: %w: %w", op, QueryErr, err)
	}
	defer rows.Close()

	messages := make(queries.OutboxDTOs, 0, q.Limit)
	for rows.Next() {
		var msgModel models.Outbox
		err = rows.Scan(
			&msgModel.ID,
			&msgModel.AggregateID,
			&msgModel.AggregateType,
			&msgModel.MessageType,
			&msgModel.MessagePayload,
			&msgModel.Status,
			&msgModel.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w: %w", op, RowScanningErr, err)
		}

		message := converters.ConvertModelToDTO(msgModel)
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (dao *OutboxDAO) UpdateStatus(ctx context.Context, c commands.OutboxCommand) error {
	const op = "postgres.OutboxDAO.UpdateStatus"

	conn := dao.Executor.GetExecutor(ctx)

	query := `UPDATE outbox SET status = $1 WHERE id = $2`

	_, err := conn.Exec(ctx, query, c.Status, c.ID)
	if err != nil {
		return fmt.Errorf("%s: %w: %w", op, StatusUpdateError, err)
	}

	return nil
}
