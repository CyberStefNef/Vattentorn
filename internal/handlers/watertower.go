package handlers

import (
	"fmt"
	"net/http"

	"github.com/CyberStefNef/Vattentorn/internal/models"
	"github.com/labstack/echo/v4"
)

var images = []models.Image{
	{URL: "/assets/images/tower1.jpg", Alt: "Bromölla", Location: "Bromölla", ID: "bromölla"},
	{URL: "/assets/images/tower2.jpg", Alt: "Kristanstad", Location: "Kristanstad", ID: "kristanstad"},
	{URL: "/assets/images/tower3.jpg", Alt: "Malmö - Hyllie", Location: "Malmö - Hyllie", ID: "malmöhyllie"},
	{URL: "/assets/images/tower4.jpg", Alt: "Malmö - Södervärn", Location: "Malmö - Södervärn", ID: "malmösödervärn"},
	{URL: "/assets/images/tower5.jpg", Alt: "Oxie", Location: "Oxie", ID: "oxie"},
	{URL: "/assets/images/tower6.jpg", Alt: "Ystad", Location: "Ystad", ID: "ystad"},
	{URL: "/assets/images/tower7.jpg", Alt: "Hälsingborg", Location: "Hälsingborg", ID: "hälsingborg"},
}

type WaterTowerPage struct {
	ID    string
	Image models.Image
}

func getWaterTower(id string) (*models.Image, error) {
	for i, image := range images {
		if id == image.ID {
			return &images[i], nil
		}
	}
	return nil, fmt.Errorf("Invalid ID!")
}

func WaterTowerHandler(c echo.Context) error {
	id := c.QueryParam("id")
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
