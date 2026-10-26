package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserModel struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	// primary key
	ID uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`

	// data
	CreatedAt time.Time `bun:"created_at,type:timestamptz,nullzero,notnull,default:current_timestamp"`
	FullName  string    `bun:"full_name,notnull"`
	UpdatedAt time.Time `bun:"updated_at,type:timestamptz,nullzero,notnull,default:current_timestamp"`
	Username  string    `bun:"username,notnull"`
}
