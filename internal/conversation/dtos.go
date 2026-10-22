package conversation

import (
	"time"

	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
)

type Conversation struct {
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	ID        uuid.UUID `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    uuid.UUID `json:"-"`
}

type ConversationMessage struct {
	CreatedAt       time.Time       `json:"createdAt"`
	Content         string          `json:"content"`
	ConversationID  uuid.UUID       `json:"conversationID"`
	ID              uuid.UUID       `json:"id"`
	Model           *string         `json:"model"`
	ParentMessageID *uuid.UUID      `json:"parentMessageID"`
	Role            schema.RoleType `json:"role"`
	SequenceNumber  int             `json:"sequenceNumber"`
	ToolCalls       []ToolCall      `json:"toolCalls"`
	Usage           *Usage          `json:"usage"`
}

type ToolCall struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type Usage struct {
	InputTokens  int `json:"inputTokens"`
	OutputTokens int `json:"outputTokens"`
	TotalTokens  int `json:"totalTokens"`
}
