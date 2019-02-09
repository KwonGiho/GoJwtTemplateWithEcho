package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"../handlers"
	"../middlewares"
	api ".."
)

func New() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	jwtGroup := e.Group("/jwt")

	middlewares.SetJwtMiddlewares(jwtGroup)

	jwtGroup.GET("/main", handlers.MainJwt)

	api.BindMainGroup(e)
	api.BindUserGroup(e)


	return e
}