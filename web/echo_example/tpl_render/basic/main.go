package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Pre-compile templates
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	// Register templates
	e.Renderer = t
	// Render a template inside your handler
	e.GET("/hello", func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "hello", "world")
	})
	e.Logger.Debug(e.Start(":80"))
}

type Template struct {
	templates *template.Template
}

// Implement echo.Renderer interface
func (t *Template) Render(w io.Writer, name string, data any, ctx echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
