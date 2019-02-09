package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func BindMainGroup(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})
}