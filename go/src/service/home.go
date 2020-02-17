package service

import (
	"net/http"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", echo.Map{"title": "Home"})
}
