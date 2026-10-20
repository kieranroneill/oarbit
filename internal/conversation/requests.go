package conversation

import "github.com/google/uuid"

type ChatRequest struct {
	ConversationID *uuid.UUID `param:"conversation_id"`
	Query          string     `json:"query"`
	UserID         uuid.UUID  `param:"user_id"`
}
