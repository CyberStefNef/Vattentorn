package handlers

import (
	"net/http"
	"strconv"

	"github.com/CyberStefNef/Vattentorn/internal/models"
	"github.com/labstack/echo/v4"
)

type ImagePage struct {
	Images      []models.Image
	CurrentPage int
	NextPage    int
	LastIndex   int
}

func ImageHandler(c echo.Context) error {
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
}
