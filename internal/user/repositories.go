package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Repository struct {
	db bun.IDB
}

/**
 * constructor functions
 */

func NewRepository(db bun.IDB) *Repository {
	return &Repository{
		db: db,
	}
}

/**
 * private functions
 */

func (r *Repository) findByID(ctx context.Context, id uuid.UUID) (*UserModel, error) {
	var model *UserModel

	err := r.db.NewSelect().
		Model(&model).
		Where("users.id = ?", id).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find user %s: %w", id, err)
	}

	return model, nil
}

/**
 * public functions
 */

func (r *Repository) FindByID(ctx context.Context, id uuid.UUID) (*UserModel, error) {
	model, err := r.findByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if model == nil {
		return nil, nil
	}

	return model, nil
}
