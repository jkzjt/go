package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/echo", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, Echo")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
