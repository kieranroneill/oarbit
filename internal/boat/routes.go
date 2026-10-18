package boat

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/uptrace/bun"

	"github.com/kieranroneill/oarbit/internal/platform/httpclient"
)

func NewGetBoatSpecificationsRoute(db *bun.DB) echo.HandlerFunc {
	return func(c *echo.Context) error {
		service := NewSpecificationService(db)
		results, err := service.FindAll(c.Request().Context())

		if err != nil {
			return err
		}

		if results == nil {
			return c.JSON(http.StatusOK, httpclient.ListResponse[Specification]{
				Items: []Specification{},
			})
		}

		return c.JSON(http.StatusOK, httpclient.ListResponse[Specification]{
			Items: results,
		})
	}
}
