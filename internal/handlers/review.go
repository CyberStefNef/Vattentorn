package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Review struct {
	User        string
	Title       string
	Description string
	Date        time.Time
	Stars       float32
}

func ReviewHandler(c echo.Context) error {
	review := Review{
		User:        "Max",
		Title:       "So toll",
		Description: "Ich liebe den Kristanstad Wasserturm",
		Date:        time.Now(),
		Stars:       1.5,
	}
	return c.Render(http.StatusOK, "review.html", review)
}
