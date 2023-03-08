package main

import "github.com/labstack/echo/v4"

func main() {

	e := echo.New()

	e.Static("/static", "assets")

	e.Logger.Debug(e.Start(":80"))

}
