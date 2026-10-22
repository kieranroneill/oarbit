package conversation

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/kieranroneill/oarbit/internal/platform/agent"
)

type Service struct {
	db                     *bun.DB
	chatAdapter            *agent.ChatAdapter
	conversationRepository *ConversationRepository
}

/**
 * static private functions
 */

func toMessageSchemas(messages []ConversationMessage) []*schema.Message {
	var _messages []*schema.Message

	for _, message := range messages {
		_messages = append(_messages, &schema.Message{
			Content: message.Content,
			Role:    message.Role,
		})
	}
	return _messages
}

func toConversationMessage(model ConversationMessageModel) ConversationMessage {
	dto := ConversationMessage{
		Content:         model.Content,
		ConversationID:  model.ConversationID,
		CreatedAt:       model.CreatedAt,
		ID:              model.ID,
		Model:           model.Model,
		Role:            schema.RoleType(model.Role),
		ParentMessageID: model.ParentMessageID,
		SequenceNumber:  model.SequenceNumber,
	}

	if model.InputTokens != nil || model.OutputTokens != nil || model.TotalTokens != nil {
		dto.Usage = &Usage{}
	}

	if model.InputTokens != nil {
		dto.Usage.InputTokens = *model.InputTokens
	}

	if model.OutputTokens != nil {
		dto.Usage.OutputTokens = *model.OutputTokens
	}

	if model.TotalTokens != nil {
		dto.Usage.TotalTokens = *model.TotalTokens
	}

	for _, toolCall := range model.ToolCalls {
		dto.ToolCalls = append(dto.ToolCalls, ToolCall{ID: toolCall.CallID, Name: toolCall.Name, Arguments: string(toolCall.Arguments)})
	}

	return dto
}

func toConversationMessageModel(message ConversationMessage) ConversationMessageModel {
	model := ConversationMessageModel{
		ID:              message.ID,
		ConversationID:  message.ConversationID,
		ParentMessageID: message.ParentMessageID,
		CreatedAt:       message.CreatedAt,
		Content:         message.Content,
		Model:           message.Model,
		Role:            string(message.Role),
		SequenceNumber:  message.SequenceNumber,
	}
	if message.Usage != nil {
		model.InputTokens = &message.Usage.InputTokens
		model.OutputTokens = &message.Usage.OutputTokens
		model.TotalTokens = &message.Usage.TotalTokens
	}
	for _, toolCall := range message.ToolCalls {
		model.ToolCalls = append(model.ToolCalls, &ConversationToolCallModel{
			ID:        uuid.New(),
			MessageID: message.ID,
			Arguments: json.RawMessage(toolCall.Arguments),
			CallID:    toolCall.ID,
			Name:      toolCall.Name,
		})
	}
	return model
}

/**
 * constructor functions
 */

func NewService(chatAdapter *agent.ChatAdapter, db *bun.DB) *Service {
	return &Service{
		db:                     db,
		chatAdapter:            chatAdapter,
		conversationRepository: NewConversationRepository(db),
	}
}

/**
 * public functions
 */

func (s *Service) Chat(ctx context.Context, query string, conversation ConversationModel) (*ConversationMessage, error) {
	messages := make([]ConversationMessage, 0, len(conversation.Messages))

	for _, message := range conversation.Messages {
		messages = append(messages, toConversationMessage(*message))
	}

	userMessage := ConversationMessage{
		CreatedAt:      time.Now(),
		Content:        query,
		ConversationID: conversation.ID,
		ID:             uuid.New(),
		Model:          nil,
		Role:           schema.User,
		SequenceNumber: 0,
		ToolCalls:      nil,
	}

	if len(messages) > 0 {
		lastMessage := messages[len(messages)-1]
		userMessage.ParentMessageID = &lastMessage.ID
		userMessage.SequenceNumber = lastMessage.SequenceNumber + 1
	}

	messageSchemas := toMessageSchemas(messages)
	messageSchemas = append(messageSchemas, &schema.Message{
		Content: userMessage.Content,
		Role:    userMessage.Role,
	})
	response, err := s.chatAdapter.Chat(ctx, messageSchemas)
	if err != nil {
		return nil, err
	}

	result := &ConversationMessage{
		ID:              uuid.New(),
		CreatedAt:       time.Now(),
		ConversationID:  conversation.ID,
		Content:         response.Content,
		Model:           &response.Model,
		ParentMessageID: &userMessage.ID,
		Role:            schema.Assistant,
		SequenceNumber:  userMessage.SequenceNumber + 1,
		Usage: &Usage{
			InputTokens:  response.Usage.InputTokens,
			OutputTokens: response.Usage.OutputTokens,
			TotalTokens:  response.Usage.TotalTokens,
		},
	}

	for _, toolCall := range response.ToolCalls {
		result.ToolCalls = append(result.ToolCalls, ToolCall{
			ID:        toolCall.ID,
			Name:      toolCall.Name,
			Arguments: toolCall.Arguments,
		})
	}

	conversation.UpdatedAt = result.CreatedAt

	if err = s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if err = NewConversationRepository(tx).Save(ctx, &conversation); err != nil {
			return err
		}

		repository := NewConversationMessageRepository(tx)
		userMessageModel := toConversationMessageModel(userMessage)
		assistantMessageModel := toConversationMessageModel(*result)

		if err = repository.Save(ctx, &userMessageModel); err != nil {
			return err
		}

		return repository.Save(ctx, &assistantMessageModel)
	}); err != nil {
		return nil, err
	}

	return result, nil
}
