package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	FullName  time.Time `json:"fullName"`
	ID        uuid.UUID `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username"`
}
