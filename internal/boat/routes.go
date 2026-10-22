package boat

import (
	"net/http"

	"github.com/kieranroneill/oarbit/internal/platform/application"
	"github.com/labstack/echo/v5"
	"github.com/uptrace/bun"

	_http "github.com/kieranroneill/oarbit/internal/platform/http"
)

func NewGetBoatSpecificationsRoute(db *bun.DB) echo.HandlerFunc {
	return func(c *echo.Context) error {
		limit, offset, err := _http.ParseListParams(c)

		if err != nil {
			return c.JSON(http.StatusBadRequest, _http.ErrorResponse{
				Errors: []application.OarbitError{
					application.NewOarbitError(application.NetworkRequestError, err.Error()),
				},
			})
		}

		return NewController(db).Specifications(c, limit, offset)
	}
}
