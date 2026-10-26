package authentication

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"

	"github.com/kieranroneill/oarbit/internal/platform/application"
	_http "github.com/kieranroneill/oarbit/internal/platform/http"
)

//==========================================
// /api/v*/auth/status
//==========================================

func NewGetAuthenticationStatusRoute(config application.Config, postgresClient *bun.DB, redisClient *redis.Client) echo.HandlerFunc {
	return func(c *echo.Context) error {
		controller := NewController(config, postgresClient)
		_user, err := controller.Authenticate(c, postgresClient, redisClient)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, _http.ErrorResponse{
				Errors: []application.OarbitError{
					*err,
				},
			})
		}

		return controller.SetupRequired(c, _user)
	}
}
