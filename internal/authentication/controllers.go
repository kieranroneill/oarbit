package authentication

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
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
 * public functions
 */

func (ctr *Controller) DeleteSessionCookie(c *echo.Context) {
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

func (ctr *Controller) SetSessionCookie(c *echo.Context, session Session) {
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
