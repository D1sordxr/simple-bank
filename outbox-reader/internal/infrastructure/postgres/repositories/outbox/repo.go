package outbox

import (
	"context"
	"github.com/D1sordxr/simple-banking-system/internal/domain/outbox"
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
	conn := tx.(pgx.Tx)
	query := ` TODO `

	rows, err := conn.Query(ctx, query, limit)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		// TODO: message := converters.ConvertModelToAggregate(msgModel)
		message := outbox.Aggregate{}
		messages = append(messages, message)
	}

	return nil, nil
}

func (r *Repository) MarkAsProcessed(ctx context.Context, tx interface{}, id string) error {

	// TODO: ...

	return nil
}
