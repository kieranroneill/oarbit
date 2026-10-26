package authentication

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ExpiresAt time.Time `json:"expiresAt"`
	ID        uuid.UUID `json:"id"`
	IssuedAt  time.Time `json:"issuedAt"`
	UserID    uuid.UUID `json:"userId"`
}
