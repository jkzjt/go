package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/search", func(ctx echo.Context) error {
		m := ctx.QueryParams()
		keywords := m["keywords"]

		fullPath := ctx.Path()

		return ctx.JSON(http.StatusOK, struct {
			Method   string   `json:"method"`
			Route    string   `json:"route"`
			Keywords []string `json:"keywords"`
			Message  string   `json:"message"`
		}{
			"GET",
			fullPath,
			keywords,
			"below are the items searched based on your typed keywords",
		})
	})

	e.POST("/list", func(ctx echo.Context) error {
		if err := ctx.Request().ParseForm(); err != nil {
			return err
		}
		m := ctx.Request().Form
		keywords := m["keywords"]

		filters := make(map[string]any)
		filters["ram"] = m["ram"]
		filters["cap"] = m["cap"]
		filters["color"] = m["color"]
		filters["rank"] = m["rank"]
		filters["sort"] = m["sort"]

		fullPath := ctx.Path()

		return ctx.JSON(http.StatusOK, echo.Map{
			"method":   ctx.Request().Method,
			"route":    fullPath,
			"keywords": keywords,
			"filters":  filters,
			"message":  "oh, god, there is no map for route use",
		})

	})

	e.PUT("/submit", func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("token")
		if token == "" {
			return echo.ErrBadRequest
		}

		if err := ctx.Request().ParseForm(); err != nil {
			return err
		}
		data := ctx.Request().PostForm

		return ctx.JSON(http.StatusOK, echo.Map{
			"method":  ctx.Request().Method,
			"route":   ctx.Path(),
			"data":    data,
			"message": "you have updated your account password and tel",
		})
	})

	e.DELETE("/cancle", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, echo.Map{
			"method":  ctx.Request().Method,
			"route":   ctx.Path(),
			"message": "we will do better in the future",
		})
	})

	e.Logger.Fatal(e.Start(":80"))
}
