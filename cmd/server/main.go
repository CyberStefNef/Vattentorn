package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Image struct {
	URL string
	Alt string
}

var images = []Image{
	{URL: "/assets/images/tower1.jpg", Alt: "Water Tower 1"},
	{URL: "/assets/images/tower2.jpg", Alt: "Water Tower 2"},
	{URL: "/assets/images/tower3.jpg", Alt: "Water Tower 3"},
	{URL: "/assets/images/tower4.jpg", Alt: "Water Tower 4"},
	{URL: "/assets/images/tower5.jpg", Alt: "Water Tower 5"},
	{URL: "/assets/images/tower6.jpg", Alt: "Water Tower 6"},
	{URL: "/assets/images/tower7.jpg", Alt: "Water Tower 7"},
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

	e.GET("/images", func(c echo.Context) error {
		page := c.QueryParam("page")
		pageIndex, err := strconv.Atoi(page)
		if err != nil || pageIndex < 0 {
			pageIndex = 0
		}

		pageSize := 5
		start := pageIndex * pageSize
		end := start + pageSize
		if start >= len(images) {
			return c.NoContent(http.StatusOK)
		}
		if end > len(images) {
			end = len(images)
		}

		return c.Render(http.StatusOK, "images.html", images[start:end])
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func HomeHandler(c echo.Context) error {
	initialImages := ""
	return c.Render(http.StatusOK, "index.html", initialImages)
}
