package conversation

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ConversationMessageModel struct {
	bun.BaseModel `bun:"table:conversation_messages,alias:cvnmsg"`

	// primary key
	ID uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`

	// foreign keys
	ConversationID  uuid.UUID  `bun:"conversation_id,type:uuid,notnull"`
	ParentMessageID *uuid.UUID `bun:"parent_message_id,type:uuid,nullzero"`

	// data
	CreatedAt      time.Time `bun:"created_at,type:timestamptz,nullzero,notnull,default:current_timestamp"`
	Content        string    `bun:"content,notnull"`
	Model          *string   `bun:"model"`
	Role           string    `bun:"role,notnull"`
	SequenceNumber int       `bun:"sequence_number,notnull"`
	InputTokens    *int      `bun:"input_tokens"`
	OutputTokens   *int      `bun:"output_tokens"`
	TotalTokens    *int      `bun:"total_tokens"`

	// relationships
	ToolCalls []*ConversationToolCallModel `bun:"rel:has-many,join:id=message_id"`
}

type ConversationModel struct {
	bun.BaseModel `bun:"table:conversations,alias:cvn"`

	// primary key
	ID uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`

	// foreign keys
	UserID uuid.UUID `bun:"user_id,type:uuid,notnull"`

	// data
	Active    bool      `bun:"active,notnull"`
	CreatedAt time.Time `bun:"created_at,type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,type:timestamptz,nullzero,notnull,default:current_timestamp"`

	// relationships
	Messages []*ConversationMessageModel `bun:"rel:has-many,join:id=conversation_id"`
}

type ConversationToolCallModel struct {
	bun.BaseModel `bun:"table:tool_calls,alias:tcll"`

	// primary key
	ID uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`

	// foreign keys
	MessageID uuid.UUID `bun:"message_id,type:uuid,notnull"`

	// data
	Arguments json.RawMessage `bun:"arguments,type:jsonb,nullzero"`
	CallID    string          `bun:"call_id,notnull,unique"`
	Name      string          `bun:"name,notnull"`
}
