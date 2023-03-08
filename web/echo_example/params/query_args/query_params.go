package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/search", func(ctx echo.Context) error {
		m := ctx.QueryParams()

		return ctx.JSON(http.StatusOK, echo.Map{
			"method":   ctx.Path(),
			"route":    ctx.Request().Method,
			"keywords": m,
		})
	})

	e.Logger.Fatal(e.Start(":80"))

}
