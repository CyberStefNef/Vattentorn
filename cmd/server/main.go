package main

import (
	"github.com/CyberStefNef/Vattentorn/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"net/http"
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

	e.Static("/assets", "web/assets")
	e.GET("/", HomeHandler)
	e.GET("/review", handlers.ReviewHandler)
	e.GET("/watertower", handlers.WaterTowerHandler)

	e.GET("/images", handlers.ImageHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

func HomeHandler(c echo.Context) error {
	initialImages := ""
	return c.Render(http.StatusOK, "index.html", initialImages)
}
