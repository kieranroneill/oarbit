package versions

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/kieranroneill/oarbit/internal/platform/application"
)

func NewGetVersionsRoute(applicationConfig application.Config) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.JSON(http.StatusOK, VersionResponse{
			Environment: applicationConfig.Environment,
			Title:       applicationConfig.Title,
			Version:     applicationConfig.Version,
		})
	}
}
