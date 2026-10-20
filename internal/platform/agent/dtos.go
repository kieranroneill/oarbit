package agent

import (
	"time"

	"github.com/cloudwego/eino/schema"
)

type ChatResult struct {
	CreatedAt time.Time       `json:"createdAt"`
	Content   string          `json:"content"`
	Model     string          `json:"model"`
	Role      schema.RoleType `json:"role"`
	ToolCalls []ToolCall      `json:"toolCalls"`
	Usage     Usage           `json:"usage"`
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
