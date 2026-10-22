package http

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v5"
)

// ParseListParams extracts limit and offset from query parameters.
//
// Parameters:
//   - c: The echo.Context object containing the HTTP request and response.
//
// Returns limit and offset as int64, and an error if the parameters are invalid.
func ParseListParams(c *echo.Context) (int64, int64, error) {
	limit := 100
	offset := 0

	if value := c.QueryParam("limit"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed <= 0 {
			return 0, 0, errors.New("limit must be a positive integer")
		}

		if parsed > 1000 {
			return 0, 0, errors.New("limit must be less than or equal to 1000")
		}

		limit = parsed
	}

	if value := c.QueryParam("offset"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed < 0 {
			return 0, 0, errors.New("offset must be a non-negative integer")
		}
		offset = parsed
	}

	return int64(limit), int64(offset), nil
}
