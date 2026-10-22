package boat

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/uptrace/bun"

	_http "github.com/kieranroneill/oarbit/internal/platform/http"
)

//==========================================
// /boats/speecfications
//==========================================

func NewGetBoatSpecificationsRoute(db *bun.DB) echo.HandlerFunc {
	return func(c *echo.Context) error {
		limit, offset, errs := _http.ParseListParams(c)

		if len(errs) > 0 {
			return c.JSON(http.StatusBadRequest, _http.ErrorResponse{
				Errors: errs,
			})
		}

		return NewController(db).Specifications(c, limit, offset)
	}
}
