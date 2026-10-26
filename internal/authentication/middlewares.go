package authentication

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"

	"github.com/kieranroneill/oarbit/internal/platform/application"
	_http "github.com/kieranroneill/oarbit/internal/platform/http"
)

func NewAuthenticateMiddleware(config application.Config, postgresClient *bun.DB, redisClient *redis.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
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

			if _user == nil {
				return c.NoContent(http.StatusUnauthorized)
			}

			// set the user in the context
			c.Set("user", _user)

			return next(c)
		}
	}
}
