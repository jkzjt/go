package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// logger middleware
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: `{"time": "${time_rfc3339_nano}", "remote_ip": "${remote_ip}", "host": "${host}", ` +
	// 		`"uri": "${uri}", "method": "${method}", "status": "${status}", "error": "${error}", ` +
	// 		`"latency_human": "${latency_human}"}` + "\n",
	// }))

	//

	// enable debug
	e.Debug = true

	// set logging
	// if l, ok := e.Logger.(*log.Logger); ok {
	// 	l.SetHeader("[${level}] ${time_rfc3339} | ${prefix} | ${short_file} | ${line}")
	// }

	e.GET("/welcome", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Welcome, "+ctx.QueryParam("name"))
	})

	//
	e.Logger.Debug(e.Start(":80"))
	// e.Start(":80")
}
