package http

import (
	"strconv"

	"github.com/kieranroneill/oarbit/internal/platform/application"
	"github.com/labstack/echo/v5"
)

// ParseListParams extracts limit and offset from query parameters.
//
// Parameters:
//   - c: The echo.Context object containing the HTTP request and response.
//
// Returns limit and offset as int64, and a list of errors if the parameters are invalid.
func ParseListParams(c *echo.Context) (int64, int64, []application.OarbitError) {
	errs := []application.OarbitError{}
	limit := 100
	offset := 0

	if value := c.QueryParam("limit"); value != "" {
		limit, err := strconv.Atoi(value)

		if err != nil || limit <= 0 {
			limit = 0
			errs = append(errs, application.NewOarbitError(application.NetworkRequestError, `"limit" must be a positive integer`))
		}

		if limit > 1000 {
			limit = 0
			errs = append(errs, application.NewOarbitError(application.NetworkRequestError, `"limit" must be less than or equal to 1000`))
		}
	}

	if value := c.QueryParam("offset"); value != "" {
		offset, err := strconv.Atoi(value)

		if err != nil || offset < 0 {
			limit = 0
			errs = append(errs, application.NewOarbitError(application.NetworkRequestError, `"offset" must be a non-negative integer`))
		}
	}

	return int64(limit), int64(offset), errs
}
