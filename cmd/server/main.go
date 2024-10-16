package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
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
	ID       string
	URL      string
	Alt      string
	Location string
}

var images = []Image{
	{URL: "/assets/images/tower1.jpg", Alt: "Bromölla", Location: "Bromölla", ID: "bromölla"},
	{URL: "/assets/images/tower2.jpg", Alt: "Kristanstad", Location: "Kristanstad", ID: "kristanstad"},
	{URL: "/assets/images/tower3.jpg", Alt: "Malmö - Hyllie", Location: "Malmö - Hyllie", ID: "malmöhyllie"},
	{URL: "/assets/images/tower4.jpg", Alt: "Malmö - Södervärn", Location: "Malmö - Södervärn", ID: "malmösödervärn"},
	{URL: "/assets/images/tower5.jpg", Alt: "Oxie", Location: "Oxie", ID: "oxie"},
	{URL: "/assets/images/tower6.jpg", Alt: "Ystad", Location: "Ystad", ID: "ystad"},
	{URL: "/assets/images/tower7.jpg", Alt: "Hälsingborg", Location: "Hälsingborg", ID: "hälsingborg"},
}

type ImagePage struct {
	Images      []Image
	CurrentPage int
	NextPage    int
	LastIndex   int
}

type WaterTowerPage struct {
	ID    string
	Image Image
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
	e.GET("/watertower", WaterTowerHandler)

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

		imagePage := ImagePage{
			Images:      images[start:end],
			CurrentPage: pageIndex,
			NextPage:    pageIndex + 1,
			LastIndex:   len(images[start:end]) - 1,
		}

		return c.Render(http.StatusOK, "images.html", imagePage)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func getWaterTower(id string) (*Image, error) {
	for i, image := range images {
		if id == image.ID {
			return &images[i], nil
		}
	}
	return nil, fmt.Errorf("Invalid ID!")
}

func WaterTowerHandler(c echo.Context) error {
	id := c.QueryParam("id")
	log.Println(id)
	waterTower, err := getWaterTower(id)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	waterTowerPage := WaterTowerPage{
		ID:    id,
		Image: *waterTower,
	}
	return c.Render(http.StatusOK, "watertower.html", waterTowerPage)
}

func HomeHandler(c echo.Context) error {
	initialImages := ""
	return c.Render(http.StatusOK, "index.html", initialImages)
}
