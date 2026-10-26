package authentication

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/uptrace/bun"
)

//==========================================
// /auth/status
//==========================================

func NewGetAuthenticationStatusRoute(db *bun.DB) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.JSON(http.StatusOK, StatusResponse{
			SetupRequired: true,
			User:          nil,
		})
	}
}
