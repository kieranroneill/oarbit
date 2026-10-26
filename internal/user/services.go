package user

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/uptrace/bun"
)

type Service struct {
	db *bun.DB
}

/**
 * static private functions
 */

func toUser(model UserModel) User {
	dto := User{
		Active:    model.Active,
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		FullName:  model.FullName,
		UpdatedAt: model.UpdatedAt,
		Username:  model.Username,
	}

	return dto
}

/**
 * constructor functions
 */

func NewService(db *bun.DB) *Service {
	return &Service{
		db: db,
	}
}

/**
 * public functions
 */

func (s *Service) Count(ctx context.Context) (int64, error) {
	count, err := s.db.NewSelect().
		Model((*UserModel)(nil)).
		Count(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}

	if err != nil {
		slog.Error("failed to find count users", "error", err)

		return 0, err
	}

	return count, nil
}
