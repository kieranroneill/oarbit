package agent

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"

	_os "github.com/kieranroneill/oarbit/internal/platform/os"
)

type ChatAdapter struct {
	agent   *adk.ChatModelAgent
	BaseURL string
	Model   string
}

func initializeAgent(ctx context.Context, baseURL string, model string) (*adk.ChatModelAgent, error) {
	chatModel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: baseURL,
		Model:   model,
	})
	if err != nil {
		return nil, err
	}

	agent, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        _os.GetEnvWithDefault("TITLE", "Oarbit"),
		Description: _os.GetEnvWithDefault("DESCRIPTION", "Oarbit is an onboard knowledge and maintenance copilot that centralizes technical documents, service logs, and operational notes for your vessel."),
		Instruction: GeneralInstruction,
		Model:       chatModel,
	})
	if err != nil {
		return nil, err
	}

	return agent, nil
}

func NewChatAdapter(ctx context.Context) (*ChatAdapter, error) {
	baseURL := os.Getenv("OLLAMA_BASE_URL")
	model := os.Getenv("OLLAMA_MODEL")

	if baseURL == "" || model == "" {
		return nil, errors.New("OLLAMA_MODEL and OLLAMA_BASE_URL must be set")
	}

	agent, err := initializeAgent(ctx, baseURL, model)
	if err != nil {
		return nil, err
	}

	return &ChatAdapter{
		agent:   agent,
		BaseURL: baseURL,
		Model:   model,
	}, nil
}

func (a *ChatAdapter) Chat(ctx context.Context, messages []*schema.Message) (*ChatResult, error) {
	var toolCalls []ToolCall

	runner := adk.NewRunner(ctx, adk.RunnerConfig{
		Agent:           a.agent,
		EnableStreaming: false,
	})
	result := &ChatResult{
		Model: a.Model,
		Role:  schema.Assistant,
	}

	chunk := runner.Run(ctx, messages)
	for {
		event, ok := chunk.Next()

		if !ok {
			break
		}

		if event.Err != nil {
			slog.Error("agent event failed", "err", event.Err)

			return nil, fmt.Errorf("agent event failed: %w", event.Err)
		}

		// continue until the last message
		if event.Output == nil || event.Output.MessageOutput == nil {
			continue
		}

		message, _, err := adk.GetMessage(event)
		if err != nil {
			slog.Error("failed to read agent message", "err", event.Err)

			return nil, fmt.Errorf("failed to read agent message: %w", err)
		}

		for _, toolCall := range message.ToolCalls {
			toolCalls = append(toolCalls, ToolCall{
				ID:        toolCall.ID,
				Name:      toolCall.Function.Name,
				Arguments: toolCall.Function.Arguments,
			})
		}

		result.Model = a.Model

		if message.ResponseMeta != nil && message.ResponseMeta.Usage != nil {
			result.Usage.InputTokens = message.ResponseMeta.Usage.PromptTokens
			result.Usage.OutputTokens = message.ResponseMeta.Usage.CompletionTokens
			result.Usage.TotalTokens = message.ResponseMeta.Usage.TotalTokens
		}

		if message.Content != "" {
			result.Content = message.Content
		}

		result.CreatedAt = time.Now()
	}

	result.ToolCalls = toolCalls

	return result, nil
}
