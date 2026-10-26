package application

import (
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v5"
)

//==========================================
// /app
//==========================================

func NewGetAppRoute(staticFiles fs.FS) echo.HandlerFunc {
	return func(c *echo.Context) error {
		indexFile, err := staticFiles.Open("index.html")
		if err != nil {
			return c.NoContent(404)
		}

		defer indexFile.Close()

		return c.Stream(http.StatusOK, "text/html", indexFile)
	}
}

//==========================================
// /healthcheck
//==========================================

func NewGetHealthcheckRoute() echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}
