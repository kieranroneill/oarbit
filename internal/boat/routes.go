package boat

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/uptrace/bun"

	"github.com/kieranroneill/oarbit/internal/platform/httpclient"
)

func NewGetBoatSpecificationsRoute(db *bun.DB) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := c.Request().Context()
		limit, offset, err := httpclient.ParseListParams(c)
		if err != nil {
			return err
		}

		service := NewService(db)
		results, total, err := service.FindAll(ctx, limit, offset)

		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		if results == nil {
			return c.JSON(http.StatusOK, httpclient.ListResponse[Specification]{
				Items:  []Specification{},
				Limit:  limit,
				Offset: offset,
				Total:  total,
			})
		}

		return c.JSON(http.StatusOK, httpclient.ListResponse[Specification]{
			Items:  results,
			Limit:  limit,
			Offset: offset,
			Total:  total,
		})
	}
}
