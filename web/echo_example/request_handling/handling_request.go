package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/create", func(ctx echo.Context) error {
		r := new(R)
		r.Code = "20000"
		r.Message = "success"
		r.Data = map[string]any{
			"username": "jt",
			"gender":   1,
			"age":      24,
		}

		if err := ctx.Bind(r); err != nil {
			return err
		}

		return ctx.JSON(http.StatusCreated, r)
	})

	e.Logger.Debug(e.Start(":80"))
}

type R struct {
	Code    string         `json:"code" xml:"code" form:"code" query:"code"`
	Message string         `json:"message" xml:"message" form:"message" query:"message"`
	Data    map[string]any `json:"data" xml:"data" form:"data" query:"data"`
}
