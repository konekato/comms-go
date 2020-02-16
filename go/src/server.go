package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"service"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Home
	e.GET("/", service.Home)
	e.GET("/home", service.Home)

	//

	e.Logger.Fatal(e.Start(":1323"))
}
