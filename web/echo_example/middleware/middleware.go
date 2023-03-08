package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Group level middleware
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		fmt.Printf("[username = %s && password = %s]\n", username, password)
		if username == "jt" && password == "123456" {
			return true, nil
		}
		return false, nil
	}))
	g.GET("/home", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Welcome!")
	})

	// Route level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			println("request to /user")
			return next(ctx)
		}
	}
	e.GET("/user", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "/user")
	}, track)

	e.Logger.Debug(e.Start(":80"))
}
