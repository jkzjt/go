package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/item/:id", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, echo.Map{
			"method": ctx.Request().Method,
			"route":  ctx.Path(),
			"id":     ctx.Param("id"),
		})
	})

	e.GET("/item/detail", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, echo.Map{
			"method":  ctx.Request().Method,
			"route":   ctx.Path(),
			"message": "this is a page on details of item",
		})
	})

	e.Logger.Fatal(e.Start(":80"))
}
