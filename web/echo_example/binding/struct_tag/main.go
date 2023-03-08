package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var label Label
	e.GET("/list/:keyword/:brand", func(ctx echo.Context) error {
		err := ctx.Bind(&label)

		if err != nil {
			return err
		}

		fmt.Printf("payload %+v", label)

		return ctx.String(http.StatusOK, "ok")
	})

	e.Start(":80")
}

type Label struct {
	Keyword string `param:"keyword" query:"keyword"`
	Brand   string `param:"brand" query:"brand"`
}
