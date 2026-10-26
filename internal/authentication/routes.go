package authentication

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
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

func NewGetAuthenticationStatusRoute(db *bun.DB, redisClient *redis.Client, config application.Config) echo.HandlerFunc {
	return func(c *echo.Context) error {
		cookie, err := c.Cookie(SessionCookieName)
		controller := NewController(config, db)

		if errors.Is(err, http.ErrNoCookie) {
			return controller.SetupRequired(c, nil)
		}

		if err != nil {
			slog.Debug("failed to get session cookie", "error", err)

			return controller.SetupRequired(c, nil)
		}

		sessionID, err := uuid.Parse(cookie.Value)

		if err != nil {
			slog.Error("failed to parse session cookie id", "error", err)

			controller.DeleteSessionCookie(c)

			return controller.SetupRequired(c, nil)
		}

		ctx := c.Request().Context()
		sessionRepository := NewSessionRepository(redisClient)
		session, err := sessionRepository.FindByID(ctx, sessionID)

		if err != nil {
			slog.Error("failed to get session from redis", "error", err)

			return c.JSON(http.StatusInternalServerError, _http.ErrorResponse{
				Errors: []application.OarbitError{
					application.NewOarbitError(application.InternalServerError, "failed to get session"),
				},
			})
		}

		// check if the session is valid
		if session == nil || !time.Now().Before(session.ExpiresAt) {
			controller.DeleteSessionCookie(c)

			return controller.SetupRequired(c, nil)
		}

		_user, err := user.NewService(db).FindByID(ctx, session.UserID)

		if err != nil {
			slog.Debug(fmt.Sprintf(`failed to get user "%s" from postgres`, session.UserID.String()), "error", err)

			return c.JSON(http.StatusInternalServerError, _http.ErrorResponse{
				Errors: []application.OarbitError{
					application.NewOarbitError(application.InternalServerError, "failed to get user details"),
				},
			})
		}

		// if no user is found for the session, delete the session
		if _user == nil {
			err = sessionRepository.DeleteByID(ctx, sessionID)

			if err != nil {
				slog.Debug(fmt.Sprintf(`failed to delete session "%s" from redis`, sessionID.String()), "error", err)
			}

			controller.DeleteSessionCookie(c)

			return controller.SetupRequired(c, nil)
		}

		// slide the ttl of the session a new session
		session, err = sessionRepository.SlideTimeoutByID(ctx, sessionID)

		if session == nil || err != nil {
			slog.Debug(fmt.Sprintf(`failed to slide session "%s" timeout`, sessionID.String()), "error", err)

			return c.JSON(http.StatusOK, StatusResponse{
				SetupRequired: false,
				User:          _user,
			})
		}

		// set the new session cookie
		controller.SetSessionCookie(c, *session)

		return c.JSON(http.StatusOK, StatusResponse{
			SetupRequired: false,
			User:          _user,
		})
	}
}
