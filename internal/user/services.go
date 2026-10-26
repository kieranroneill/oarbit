package user

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Service struct {
	db         *bun.DB
	repository *Repository
}

/**
 * constructor functions
 */

func NewService(db *bun.DB) *Service {
	return &Service{
		db:         db,
		repository: NewRepository(db),
	}
}

/**
 * private functions
 */

func toUser(model *UserModel) *User {
	dto := &User{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		FullName:  model.FullName,
		UpdatedAt: model.UpdatedAt,
		Username:  model.Username,
	}

	return dto
}

/**
 * public functions
 */

func (s *Service) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	model, err := s.repository.FindByID(ctx, id)

	if err != nil {
		slog.Error("failed to find user by id", "error", err)

		return nil, err
	}

	if model == nil {
		return nil, nil
	}

	return toUser(model), nil
}

func (s *Service) Count(ctx context.Context) (int64, error) {
	count, err := s.db.NewSelect().
		Model((*UserModel)(nil)).
		Count(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}

	if err != nil {
		slog.Error("failed to count users", "error", err)

		return 0, err
	}

	return count, nil
}
