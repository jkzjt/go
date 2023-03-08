package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	t := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/e", func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "template.html", map[string]any{
			"name": "Dolly!",
		})
	}).Name = "foobar"

	e.Logger.Debug(e.Start(":80"))
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data any, ctx echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]any); isMap {
		viewContext["reverse"] = ctx.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
