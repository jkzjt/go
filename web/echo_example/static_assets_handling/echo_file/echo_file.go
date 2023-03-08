package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	e.File("/", "public/index.html")

	e.Logger.Debug(e.Start(":80"))
}
