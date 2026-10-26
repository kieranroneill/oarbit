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

type Controller struct {
	applicationConfig application.Config
	userService       *user.Service
}

/**
 * constructor functions
 */

func NewController(applicationConfig application.Config, db *bun.DB) *Controller {
	return &Controller{
		applicationConfig: applicationConfig,
		userService:       user.NewService(db),
	}
}

/**
 * private functions
 */

func (ctr *Controller) deleteSessionCookie(c *echo.Context) {
	c.SetCookie(&http.Cookie{
		Expires:  time.Unix(1, 0),
		Name:     SessionCookieName,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   ctr.applicationConfig.Environment != "development",
		SameSite: http.SameSiteLaxMode,
		Value:    "",
	})
}

func (ctr *Controller) setSessionCookie(c *echo.Context, session Session) {
	c.SetCookie(&http.Cookie{
		Expires:  session.ExpiresAt,
		Name:     SessionCookieName,
		Path:     "/",
		HttpOnly: true,
		Secure:   ctr.applicationConfig.Environment != "development",
		SameSite: http.SameSiteLaxMode,
		Value:    session.ID.String(),
	})
}

/**
 * public functions
 */

func (ctr *Controller) Authenticate(c *echo.Context, postgresClient *bun.DB, redisClient *redis.Client) (*user.User, *application.OarbitError) {
	cookie, err := c.Cookie(SessionCookieName)

	if errors.Is(err, http.ErrNoCookie) {
		return nil, nil
	}

	if err != nil {
		slog.Debug("failed to get session cookie", "error", err)

		return nil, nil
	}

	sessionID, err := uuid.Parse(cookie.Value)

	if err != nil {
		slog.Error("failed to parse session cookie id", "error", err)

		ctr.deleteSessionCookie(c)

		return nil, nil
	}

	ctx := c.Request().Context()
	sessionRepository := NewSessionRepository(redisClient)
	session, err := sessionRepository.FindByID(ctx, sessionID)

	if err != nil {
		slog.Error("failed to get session from redis", "error", err)

		_error := application.NewOarbitError(application.InternalServerError, "failed to get session")

		return nil, &_error
	}

	// check if the session is valid
	if session == nil || !time.Now().Before(session.ExpiresAt) {
		ctr.deleteSessionCookie(c)

		return nil, nil
	}

	_user, err := user.NewService(postgresClient).FindByID(ctx, session.UserID)

	if err != nil {
		slog.Debug(fmt.Sprintf(`failed to get user "%s" from postgres`, session.UserID.String()), "error", err)

		_error := application.NewOarbitError(application.InternalServerError, "failed to get user details")

		return nil, &_error
	}

	// if no user is found for the session, delete the session
	if _user == nil {
		err = sessionRepository.DeleteByID(ctx, sessionID)

		if err != nil {
			slog.Debug(fmt.Sprintf(`failed to delete session "%s" from redis`, sessionID.String()), "error", err)
		}

		ctr.deleteSessionCookie(c)

		return nil, nil
	}

	// slide the ttl of the session a new session
	session, err = sessionRepository.SlideTimeoutByID(ctx, sessionID)

	if session == nil || err != nil {
		slog.Debug(fmt.Sprintf(`failed to slide session "%s" timeout`, sessionID.String()), "error", err)

		return _user, nil
	}

	// set the new session cookie
	ctr.setSessionCookie(c, *session)

	return _user, nil
}

func (ctr *Controller) SetupRequired(c *echo.Context, _user *user.User) error {
	ctx := c.Request().Context()
	users, err := ctr.userService.Count(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _http.ErrorResponse{
			Errors: []application.OarbitError{
				application.NewOarbitError(application.InternalServerError, "failed to fetch get authentication status"),
			},
		})
	}

	if users > 0 {
		return c.JSON(http.StatusOK, StatusResponse{
			SetupRequired: false,
			User:          _user,
		})
	}

	return c.JSON(http.StatusOK, StatusResponse{
		SetupRequired: true,
		User:          _user,
	})
}
