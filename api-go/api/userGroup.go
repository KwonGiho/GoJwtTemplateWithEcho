package api

import (
	"github.com/labstack/echo"
	"./handlers"
)

func BindUserGroup(e *echo.Echo) {
	e.GET("/login", handlers.Login)
}