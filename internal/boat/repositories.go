package boat

import "github.com/uptrace/bun"

/**
 * profile
 */

type ProfileRepository struct {
	db *bun.DB
}

func NewProfileRepository(db *bun.DB) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

/**
 * specification
 */

type SpecificationRepository struct {
	db *bun.DB
}

func NewSpecificationRepository(db *bun.DB) *SpecificationRepository {
	return &SpecificationRepository{
		db: db,
	}
}
