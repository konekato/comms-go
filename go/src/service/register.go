package service

import (
	"net/http"

	"github.com/labstack/echo"
)

func RegisterForm(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", echo.Map{"title": "Register"})
}

func Register(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
