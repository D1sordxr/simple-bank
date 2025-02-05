package outbox

import (
	"context"
	"errors"
	"github.com/D1sordxr/simple-banking-system/internal/domain/outbox"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/converters"
	"github.com/D1sordxr/simple-banking-system/internal/infrastructure/postgres/models"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn *pgx.Conn
}

func NewOutboxRepository(conn *pgx.Conn) *Repository {
	return &Repository{Conn: conn}
}

func (r *Repository) FetchPendingMessages(ctx context.Context, tx interface{}, limit int) ([]outbox.Aggregate, error) {
	conn, ok := tx.(pgx.Tx)
	if !ok {
		return nil, InvalidTxType
	}

	query := `SELECT id, aggregate_id, aggregate_type, message_type, message_payload, status, created_at
				FROM outbox
				WHERE status = 'pending'
				ORDER BY created_at
				LIMIT $1`

	rows, err := conn.Query(ctx, query, limit)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []outbox.Aggregate{}, nil
		}
		return nil, QueryErr
	}
	defer rows.Close()

	messages := make([]outbox.Aggregate, 0, limit)
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
			return nil, RowScanningErr
		}
		message := converters.ConvertModelToAggregate(msgModel)
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *Repository) MarkAsProcessed(ctx context.Context, tx interface{}, id string) error {

	// TODO: ...

	return nil
}
