package vessel

import (
	"net/http"

	"github.com/kieranroneill/oarbit/internal/platform/application"
	"github.com/labstack/echo/v5"
	"github.com/uptrace/bun"

	_http "github.com/kieranroneill/oarbit/internal/platform/http"
)

type Controller struct {
	service *Service
}

/**
 * constructor functions
 */

func NewController(db *bun.DB) *Controller {
	return &Controller{
		service: NewService(db),
	}
}

/**
 * public functions
 */

func (ctr *Controller) Specifications(c *echo.Context, limit int64, offset int64) error {
	ctx := c.Request().Context()
	results, total, err := ctr.service.FindSpecifications(ctx, limit, offset)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _http.ErrorResponse{
			Errors: []application.OarbitError{
				application.NewOarbitError(application.InternalServerError, "failed to fetch boat specifications"),
			},
		})
	}

	if results == nil {
		return c.JSON(http.StatusOK, _http.ListResponse[Specification]{
			Items:  []Specification{},
			Limit:  limit,
			Offset: offset,
			Total:  total,
		})
	}

	return c.JSON(http.StatusOK, _http.ListResponse[Specification]{
		Items:  results,
		Limit:  limit,
		Offset: offset,
		Total:  total,
	})
}
