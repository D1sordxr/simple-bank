package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/commands"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/application/queries"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres/converters"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/postgres/models"
	"github.com/jackc/pgx/v5"
)

type OutboxDAO struct {
	*postgres.Connection
}

func NewOutboxDAO(connection *postgres.Connection) *OutboxDAO {
	return &OutboxDAO{Connection: connection}
}

func (dao *OutboxDAO) FetchMessages(ctx context.Context, q queries.OutboxQuery) (queries.OutboxDTOs, error) {
	const op = "postgres.OutboxDAO.FetchMessages"

	// TODO: if tx { conn := tx }

	query := `SELECT id, aggregate_id, aggregate_type, message_type, message_payload, status, created_at
				FROM outbox
				WHERE status = $1
				ORDER BY created_at
				LIMIT $2`

	rows, err := dao.Connection.Query(ctx, query, q.Status, q.Limit)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return queries.OutboxDTOs{}, nil
		}
		//return nil, fmt.Errorf("%s: %w", op, QueryErr)
		return nil, fmt.Errorf("%s: %w", op, err)
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
			return nil, fmt.Errorf("%s: %w", op, RowScanningErr)
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

	// TODO: if tx { conn := tx }

	query := `UPDATE outbox SET status = $1 WHERE id = $2`

	_, err := dao.Connection.Exec(ctx, query, c.Status, c.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, StatusUpdateError)
	}

	return nil
}
