package conversation

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/uptrace/bun"

	"github.com/kieranroneill/oarbit/internal/platform/agent"
)

func NewPostChatRoute(chatAdapter *agent.ChatAdapter, db *bun.DB) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var params ChatRequest

		controller := NewController(chatAdapter, db)

		if err := c.Bind(&params); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if params.Query == "" {
			return echo.NewHTTPError(http.StatusBadRequest, `"query"" is required`)
		}

		if params.ConversationID != nil {
			return controller.ChatWithConversationID(c, params.Query, params.UserID, *params.ConversationID)
		}

		return controller.Chat(c, params.Query, params.UserID)
	}
}
