package conversation

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

/**
 * conversation
 */

type ConversationRepository struct {
	db bun.IDB
}

func NewConversationRepository(db bun.IDB) *ConversationRepository {
	return &ConversationRepository{
		db: db,
	}
}

func (r *ConversationRepository) FindByIDAndUserIDWithMessages(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*ConversationModel, error) {
	conversation := new(ConversationModel)
	err := r.db.NewSelect().
		Model(conversation).
		Where("cvn.id = ?", id).
		Where("cvn.user_id = ?", userID).
		Relation("Messages", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("cvnmsg.sequence_number ASC")
		}).
		Relation("Messages.ToolCalls").
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		slog.Error("failed to find conversation by id and user id", "error", err)

		return nil, err
	}

	return conversation, nil
}

// Save inserts a conversation or updates its mutable fields when its ID already exists.
func (r *ConversationRepository) Save(ctx context.Context, conversation *ConversationModel) error {
	_, err := r.db.NewInsert().
		Model(conversation).
		On("CONFLICT (id) DO UPDATE").
		Set("active = EXCLUDED.active").
		Set("updated_at = EXCLUDED.updated_at").
		Exec(ctx)

	if err != nil {
		slog.Error("failed to save conversation", "error", err)

		return err
	}

	return nil
}

/**
 * conversation message
 */

type ConversationMessageRepository struct {
	db bun.IDB
}

func NewConversationMessageRepository(db bun.IDB) *ConversationMessageRepository {
	return &ConversationMessageRepository{
		db: db,
	}
}

// Save inserts a message and its tool calls. Use a transaction to save them atomically.
func (r *ConversationMessageRepository) Save(ctx context.Context, message *ConversationMessageModel) error {
	if _, err := r.db.NewInsert().Model(message).Exec(ctx); err != nil {
		slog.Error("failed to save conversation message", "error", err)

		return err
	}

	if len(message.ToolCalls) > 0 {
		_, err := r.db.NewInsert().Model(&message.ToolCalls).Exec(ctx)

		if err != nil {
			slog.Error("failed to save conversation message tool calls", "error", err)

			return err
		}
	}

	return nil
}
