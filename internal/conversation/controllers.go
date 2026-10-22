package conversation

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/uptrace/bun"

	"github.com/kieranroneill/oarbit/internal/platform/agent"
)

type Controller struct {
	conversationRepository *ConversationRepository
	service                *Service
}

/**
 * constructor functions
 */

func NewController(chatAdapter *agent.ChatAdapter, db *bun.DB) *Controller {
	return &Controller{
		conversationRepository: NewConversationRepository(db),
		service:                NewService(chatAdapter, db),
	}
}

/**
 * public functions
 */

func (ctr *Controller) Chat(c *echo.Context, query string, userID uuid.UUID) error {
	ctx := c.Request().Context()
	conversation := ConversationModel{
		Active:    true,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Messages:  []*ConversationMessageModel{},
		UpdatedAt: time.Now(),
		UserID:    userID,
	}

	result, err := ctr.service.Chat(ctx, query, conversation)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, result)
}

func (ctr *Controller) ChatWithConversationID(c *echo.Context, query string, userID uuid.UUID, conversationID uuid.UUID) error {
	ctx := c.Request().Context()
	conversation, err := ctr.service.conversationRepository.FindByIDAndUserIDWithMessages(ctx, conversationID, userID)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if conversation == nil {
		return echo.NewHTTPError(http.StatusNotFound, "conversation not found")
	}

	result, err := ctr.service.Chat(ctx, query, *conversation)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, result)
}
