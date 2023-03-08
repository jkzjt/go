package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// default
	// e.Use(middleware.Static("assets")) // "/assets" stands for context root starting
	// custom configuration
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "assets",
		Browse: true,
	}))
	// Default behavior when using with non root URL paths is to append the URL path to the filesystem path.
	// g := e.Group("images")
	// g.Use(middleware.Static("assets"))
	e.Logger.Debug(e.Start(":80"))
}
