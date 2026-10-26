package authentication

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"

	"github.com/kieranroneill/oarbit/internal/platform/application"
	_http "github.com/kieranroneill/oarbit/internal/platform/http"
	"github.com/kieranroneill/oarbit/internal/user"
)

//==========================================
// /auth/status
//==========================================

func NewGetAuthenticationStatusRoute(db *bun.DB, redisClient *redis.Client) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := c.Request().Context()
		users, err := user.NewService(db).Count(ctx)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, _http.ErrorResponse{
				Errors: []application.OarbitError{
					application.NewOarbitError(application.InternalServerError, "failed to fetch get authentication status"),
				},
			})
		}

		// if setup is not required, attempt to get user from cookie
		if users > 0 {
			// TODO: get user from cookie
			return c.JSON(http.StatusOK, StatusResponse{
				SetupRequired: false,
				User:          nil,
			})
		}

		return c.JSON(http.StatusOK, StatusResponse{
			SetupRequired: true,
			User:          nil,
		})
	}
}
