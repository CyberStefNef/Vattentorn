package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	renderer := &Template{
		templates: template.Must(template.ParseGlob("web/templates/*.html")),
	}
	e.Renderer = renderer

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "web/assets") // Serve static assets
	e.GET("/", HomeHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

func HomeHandler(c echo.Context) error {
	data := map[string]interface{}{
		"Title": "Welcome to My Project",
	}
	return c.Render(http.StatusOK, "index.html", data)
}
