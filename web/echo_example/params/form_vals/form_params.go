package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/login", func(ctx echo.Context) error {
		username := ctx.FormValue("username")
		password := ctx.FormValue("password")

		return ctx.JSON(http.StatusOK, echo.Map{
			"method":   ctx.Request().Method,
			"route":    ctx.Path(),
			"username": username,
			"password": password,
		})
	})

	e.Logger.Fatal(e.Start(":80"))
}
